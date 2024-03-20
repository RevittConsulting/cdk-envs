import { TxFormData } from "@/components/tx/tx-form";
import { ApiResponse } from "../types";
import { TxResponse } from "@/types/tx";

export interface TxApi {
  submitTx: (req: TxFormData) => Promise<ApiResponse<TxResponse>>;
}