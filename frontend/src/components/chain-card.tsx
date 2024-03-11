"use client";

import { useState, useEffect, use } from "react";
import { Chain, ChainData } from "@/types/chain";
import { Button } from "@/components/ui/button";
import { restartServices, stopAllServices } from "@/api/chain";
import { PlayIcon, StopIcon } from "@heroicons/react/16/solid";
import { Link } from "lucide-react";
import { useChainContext } from "@/context/chain-context";
import Spinner from "@/components/spinner";

export default function ChainCard({
  chain,
  data,
}: {
  chain: Chain;
  data: ChainData;
}) {
  const { activeChain, setActiveChain } = useChainContext();
  const [dataDisplay, setDataDisplay] = useState<ChainData>({} as ChainData);

  useEffect(() => {
    console.log("ChainCard data", data);
    if (activeChain === chain.serviceName) {
      setDataDisplay(data);
    }
  }, [data, activeChain, chain.serviceName]);

  const startServices = async () => {
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
      <div className="p-4 rounded-lg border w-[36vw]">
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
            <p className="font-thin">
              {dataDisplay.mostRecentL1Block
                ? dataDisplay.mostRecentL1Block
                : 0}
            </p>
            {renderL1ClickableLinks(`block/${chain.L1.latestL1BlockNumber}`)}
          </div>

          <hr />
          <div className="flex justify-between items-center">
            <h3 className="text-sm font-semibold pb-2">L2</h3>
            <h3 className="text-sm font-semibold pb-2">
              Chain ID: <span className="font-normal">{chain.L2.chainId}</span>
            </h3>
          </div>
          <div className="flex gap-2">
            <p>Datastreamer URL</p>
            <p className="font-thin">{chain.L2.datastreamerUrl}</p>
          </div>
          <div>
            <p>Latest L2 Batch Number</p>
            <p className="font-thin">
              {dataDisplay.mostRecentL2Batch
                ? dataDisplay.mostRecentL2Batch
                : 0}
            </p>
            {renderL2ClickableLinks(`batch/${chain.L2.latestBatchNumber}`)}
          </div>
        </div>
      </div>
    )
  );
}
