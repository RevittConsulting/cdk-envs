"use client";

import { useState, useEffect, useRef } from "react";
import ChainCard from "@/components/chain-card";
import { Chain } from "../types/chain";
import { getChains } from "@/api/chain";

export default function Home() {
  const [chains, setChains] = useState<Chain[]>([]);

  useEffect(() => {
    const fetchChains = async () => {
      const c = await getChains();
      console.log(c);
      setChains(c);
    };

    fetchChains();
  }, []);

  return (
    <div className="py-20 mx-20">
        <div className="flex w-max space-x-4 p-4">
          {chains.map((chain, index) => (
            <ChainCard key={index} chain={chain} />
          ))}
        </div>
    </div>
  );
}
