"use client";

import React, { useState, useEffect } from "react";
import axios from "axios";
import { Chain } from "../types/chain";

export default function ChainCard() {
  const [chainData, setChainData] = useState<Chain | null>(null);

  const getChainData = async () => {
    try {
      const response = await axios.get<Chain>(
        "http://localhost:80/api/v1/chain"
      );
      setChainData(response.data);
    } catch (error) {
      console.error("Error fetching chain data:", error);
    }
  };

  useEffect(() => {
    getChainData();
  }, []);

  const renderKeyValuePairs = (data: { [key: string]: string | number }) => (
    <div className="p-4">
      {Object.entries(data).map(([key, value]) => (
        <p key={key} className="mb-2 mx-2">
          <strong>{key}:</strong>{" "}
          <span className="dark:text-gray-400 text-gray-600">{value}</span>
        </p>
      ))}
    </div>
  );

  return (
    chainData && (
      <div className="p-4 rounded-lg shadow-md border">
        <div className="flex sm:max-h-[20vh] lg:max-h-[50vh] max-h-[50vh] overflow-y-auto overflow-x-auto">
          <div className="">
            <h2 className="text-lg font-semibold">L1</h2>
            {renderKeyValuePairs(chainData.L1)}
          </div>
          <div className="mt-4">
            <h2 className="text-lg font-semibold">L2</h2>
            {renderKeyValuePairs(chainData.L2)}
          </div>
        </div>
      </div>
    )
  );
}
