"use client";

import { useState, useEffect, use } from "react";
import { Chain, ChainData } from "@/types/chain";
import { Button } from "@/components/ui/button";
import { restartServices, stopAllServices } from "@/api/chain";
import { PlayIcon, StopIcon } from "@heroicons/react/16/solid";
import { Link } from "lucide-react";
import { useChainContext } from "@/context/chain-context";
import Spinner from "@/components/spinner";
import TripleDotLoader from "@/components/dot-ellipsis";

export default function ChainCard({
  chain,
  data,
}: {
  chain: Chain;
  data: ChainData;
}) {
  const { activeChain, setActiveChain } = useChainContext();
  const [mostRecentL1Block, setMostRecentL1Block] = useState<number>(0);
  const [mostRecentL2Batch, setMostRecentL2Batch] = useState<number>(0);

  useEffect(() => {
    if (activeChain === chain.serviceName) {
      setMostRecentL1Block(data.mostRecentL1Block);
      setMostRecentL2Batch(data.mostRecentL2Batch);
    }
  }, [data]);

  const startServices = async () => {
    await stopAllServices();
    console.log("Starting RPC services for", chain.serviceName);
    const response = await restartServices(chain.serviceName);
    if (response.status >= 200 && response.status < 300) {
      setActiveChain(chain.serviceName);
    }
    console.log(response);
  };

  const stopServices = async () => {
    console.log("Stopping RPC services for", chain.serviceName);
    const response = await stopAllServices();
    if (response.status >= 200 && response.status < 300) {
      setActiveChain(null);
    }
    console.log(response);
  };

  const renderL1ClickableLinks = (url: string) => {
    return (
      <div className="flex gap-2 mt-1">
        <a
          href={`https://sepolia.etherscan.io/${url}`}
          target="_blank"
          className="font-thin flex gap-2 bg-accent-foreground/20 hover:bg-accent-foreground/30 px-1 rounded-md"
        >
          etherscan
          <span>
            <Link width={14} />
          </span>
        </a>
        <a
          href={`https://zkevm.blockscout.com/${url}`}
          target="_blank"
          className="font-thin flex gap-2 bg-accent-foreground/20 hover:bg-accent-foreground/30 px-1 rounded-md"
        >
          blockscout
          <span>
            <Link width={14} />
          </span>
        </a>
      </div>
    );
  };

  const renderL2ClickableLinks = (url: string) => {
    return (
      <div className="flex gap-2 mt-1">
        <a
          href={`https://zkevm.polygonscan.com/${url}`}
          target="_blank"
          className="font-thin flex gap-2 bg-accent-foreground/20 hover:bg-accent-foreground/30 px-1 rounded-md"
        >
          polygonscan
          <span>
            <Link width={14} />
          </span>
        </a>
      </div>
    );
  };

  return (
    chain && (
      <div
        className={`p-4 rounded-lg border w-[30vw] ${
          activeChain === chain.serviceName
            ? "border-purple-500 ring-2 ring-purple-500 ring-opacity-50"
            : ""
        }`}
      >
        <div className="flex flex-col gap-2">
          <div className="flex items-center justify-between">
            <h1 className="text-lg font-semibold">{chain.networkName}</h1>
            <span className="text-primary">(Sepolia)</span>
          </div>

          <div className="flex w-full justify-center items-center gap-4">
            <Button
              onClick={startServices}
              className="w-full"
              disabled={activeChain === chain.serviceName}
            >
              {activeChain === chain.serviceName ? (
                <>
                  <span className="mr-1">
                    <Spinner />
                  </span>
                  RPC Services Running
                </>
              ) : (
                <>
                  <span className="mr-1">
                    <PlayIcon className="w-5 h-5" />
                  </span>
                  Start RPC Services
                </>
              )}
            </Button>
            <Button
              onClick={stopServices}
              variant={"outlineprimary"}
              className="w-full"
              disabled={activeChain !== chain.serviceName}
            >
              <span className="mr-1">
                <StopIcon className="w-5 h-5" />
              </span>
              Stop RPC Services
            </Button>
          </div>

          <hr />

          <div className="flex justify-between items-center">
            <h2 className="text-sm font-semibold pb-2">L1</h2>
            <h2 className="text-sm font-semibold pb-2">
              Chain ID: <span className="font-normal">{chain.L1.chainId}</span>
            </h2>
          </div>
          <div className="flex gap-2">
            <p>Rpc URL</p>
            <p className="font-thin">{chain.L1.rpcUrl}</p>
          </div>
          <div>
            <p>Rollup Manager Address</p>
            <p className="font-thin">{chain.L1.rollupManagerAddress}</p>
            {renderL1ClickableLinks(`address/${chain.L1.rollupManagerAddress}`)}
          </div>
          <div>
            <p>Rollup Address</p>
            <p className="font-thin">{chain.L1.rollupAddress}</p>
            {renderL1ClickableLinks(`address/${chain.L1.rollupAddress}`)}
          </div>
          <div>
            <p>Latest L1 Block</p>
            <div className="flex items-center gap-2">
              <p className="font-thin">{mostRecentL1Block}</p>
              {activeChain === chain.serviceName && <TripleDotLoader />}
            </div>
            {renderL1ClickableLinks(`block/${mostRecentL1Block}`)}
          </div>

          <hr />
          <div className="flex justify-between items-center">
            <h3 className="text-sm font-semibold pb-2">L2</h3>
            <h3 className="text-sm font-semibold pb-2">
              Chain ID: <span className="font-normal">{chain.L2.chainId}</span>
            </h3>
          </div>
          <div className="">
            <p>Datastreamer URL</p>
            <p className="font-thin">{chain.L2.datastreamerUrl}</p>
          </div>
          <div>
            <p>Latest L2 Batch Number</p>
            <div className="flex items-center gap-2">
              <p className="font-thin">{mostRecentL2Batch}</p>
              {activeChain === chain.serviceName && <TripleDotLoader />}
            </div>
            {renderL2ClickableLinks(`batch/${mostRecentL2Batch}`)}
          </div>
        </div>
      </div>
    )
  );
}
