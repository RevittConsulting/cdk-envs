export type Chain = {
  L1: L1;
  L2: L2;
};

type L1 = {
  [key: string]: string | number;
};

type L2 = {
  [key: string]: string | number;
};