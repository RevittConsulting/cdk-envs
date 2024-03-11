"use client";

import { useState, useEffect, useRef } from "react";
import ChainCard from "@/components/chain-card";
import { Chain, ChainData } from "../types/chain";
import { getChains } from "@/api/chain";

export default function Home() {
  const [chains, setChains] = useState<Chain[]>([]);

  const [data, setData] = useState<ChainData>({} as ChainData);
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/api/v1/ws");

    ws.onopen = () => {
      console.log("Connected to the server");
    };

    ws.onmessage = (event) => {
      console.log("Received data from the server", event.data);
      const newData = event.data;
      setData(newData);
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

  useEffect(() => {
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
          {chains.map((chain, index) => (
            <ChainCard key={index} chain={chain} data={data} />
          ))}
        </div>
    </div>
  );
}
