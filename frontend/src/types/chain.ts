export type Chain = {
  networkName: string;
  L1: L1;
  L2: L2;
  lastUpdated: Date;
  serviceName: string;
};

type L1 = {
  chainId: string;
  rpcUrl: string;
  rollupManagerAddress: string;
  rollupAddress: string;

  latestL1BlockNumber: number;
  highestSequencedBatch: number;
  highestVerifiedBatch: number;
};

type L2 = {
  chainId: string;
  datastreamerUrl: string;

  latestBatchNumber: number;
  latestBlockNumber: number;
  datastreamerStatus: string;
};

export type ChainData = {
  mostRecentL1Block: number;
  mostRecentL2Batch: number;
}