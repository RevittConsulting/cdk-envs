import { TxResponse } from "@/types/tx";
import { TxFormData } from "@/components/tx/tx-form";
import { TxApi } from './interface'
import { http } from '../http'

const Tx_Url = '/tx'

export const createTxApi = (): TxApi => ({
  submitTx: async (req: TxFormData) => await http.post<TxResponse>(Tx_Url, req)
})