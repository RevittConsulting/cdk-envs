"use client";

import { useState, useEffect, useRef } from "react";
import ChainCard from "@/components/chain-card";
import { Chain, ChainData } from "../types/chain";
import { getChains, stopAllServices } from "@/api/chain";
import { useChainContext } from "@/context/chain-context";

export default function Home() {
  const [chains, setChains] = useState<Chain[]>([]);
  const { activeChain, setActiveChain } = useChainContext();

  const [data, setData] = useState<ChainData>({} as ChainData);
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/api/v1/ws");

    ws.onopen = () => {
      console.log("Connected to the server");
    };

    ws.onmessage = (event) => {
      console.log("Received data from the server", event.data);
      try {
        const newData: ChainData = JSON.parse(event.data);
        setData(newData);
      } catch (error) {
        console.error("Error parsing data:", error);
      }
    };

    ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    ws.onclose = () => {
      console.log("Disconnected from the server");
    };

    return () => {
      ws.close();
    };
  }, []);

  const cleanup = async () => {
    const res = await stopAllServices();
    if (res.status >= 200 && res.status < 300) {
      setActiveChain(null);
    }
  };

  useEffect(() => {
    cleanup();

    const fetchChains = async () => {
      const res = await getChains();
      if (res.status >= 200 && res.status < 300) {
        setChains(res.data);
      } else {
        console.error("Error fetching chains");
      }
    };

    fetchChains();
  }, []);

  return (
    <div className="mt-4 mx-20">
      <div className="flex gap-4 justify-center p-4">
        {chains.length > 0 ? (
          chains.map((chain, index) => (
            <ChainCard key={index} chain={chain} data={data} />
          ))
        ) : (
          <div>No chains found</div>
        )}
      </div>
    </div>
  );
}
