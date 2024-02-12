export type Chain = {
  id: number;
  networkName: string;
  L1: L1;
  L2: L2;
  lastUpdated: Date;
};

type L1 = {
  [key: string]: string | number;
};

type L2 = {
  [key: string]: string | number;
};