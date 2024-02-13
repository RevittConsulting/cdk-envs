"use client";

import { Chain } from "@/types/chain";

export default function ChainCard({ chain }: { chain: Chain }) {

  const renderKeyValuePairs = (data: { [key: string]: string | number }) => (
    <div>
      {Object.entries(data).map(([key, value]) => (
        <p key={key}>
          <strong className="text-sm">{key}:</strong>{" "}
          <span className="dark:text-gray-400 text-gray-600 text-sm">{value}</span>
        </p>
      ))}
    </div>
  );

  return (
    chain && (
      <div className="p-4 rounded-lg border w-[26vw]">
        <div className="flex flex-col gap-4">
          <div className="flex items-center justify-between">
            <h1 className="text-lg font-semibold">{chain.networkName}</h1> <span className="text-primary">(Sepolia)</span>
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
