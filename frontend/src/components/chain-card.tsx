"use client";

import { useState, useEffect } from "react";
import { Chain } from "@/types/chain";
import { Button } from "@/components/ui/button";
import { restartServices } from "@/api/chain";

export default function ChainCard({ chain }: { chain: Chain }) {
  const [data, setData] = useState<any>(null);
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/api/v1/ws");

    ws.onopen = () => {
      console.log('Connected to the server');
    };

    ws.onmessage = (event) => {
      const newData = event.data;
      setData(newData);
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.onclose = () => {
      console.log('Disconnected from the server');
    };

    return () => {
      ws.close();
    };
  }, []);

  const startServices = async () => {
    console.log("Starting RPC services for", chain.serviceName);
    const response = await restartServices(chain.serviceName);
    console.log(response);
  };

  const renderKeyValuePairs = (data: { [key: string]: string | number }) => (
    <div>
      {Object.entries(data).map(([key, value]) => (
        <p key={key}>
          <strong className="text-sm">{key}:</strong>{" "}
          <span className="dark:text-gray-400 text-gray-600 text-sm">
            {value}
          </span>
        </p>
      ))}
    </div>
  );

  return (
    chain && (
      <div className="p-4 rounded-lg border w-[30vw]">
        <div className="flex flex-col gap-4">
          <Button onClick={startServices}>Start RPC Services</Button>
          <h1 className="text-lg font-semibold">{data}</h1>
          <div className="flex items-center justify-between">
            <h1 className="text-lg font-semibold">{chain.networkName}</h1>
            <span className="text-primary">(Sepolia)</span>
          </div>
          <div>
            <h2 className="text-sm font-semibold pb-2">L1</h2>
            {renderKeyValuePairs(chain.L1)}
          </div>
          <div>
            <h3 className="text-sm font-semibold pb-2">L2</h3>
            {renderKeyValuePairs(chain.L2)}
          </div>
        </div>
      </div>
    )
  );
}
