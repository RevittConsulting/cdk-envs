'use client'

import { useState, useEffect } from "react";
import ChainCard from "@/components/chain-card";
import { Chain } from "../types/chain";
import { getChains } from "@/api/chain"

export default function Home() {
  const [chains, setChains] = useState<Chain[]>([]);

  useEffect(() => {
    const fetchChains = async () => {
      const c = await getChains();
      console.log(c)
      setChains(c);
    };
    
    fetchChains();
  }, []);

  return (
    <div className="py-10">
    <div className="flex flex-row items-center justify-center gap-4 flex-wrap">
      {chains.map((chain, index) => (
        <ChainCard key={index} chain={chain} />
      ))}
    </div>
    </div>
  );
}
