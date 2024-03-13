export type TxOutput = {
  signedTx: string;
  fromAddress: string;
  balance: number;
};

export type TxResponse = {
  output: TxOutput;
  error: string;
};