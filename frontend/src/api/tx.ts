import { axiosInstance } from '@/utils/axios-instance'
import { AxiosResponse } from 'axios';
import { TxResponse } from "@/types/tx";
import { TxFormData } from "@/components/tx/tx-form";

export const submitTx = async (req: TxFormData): Promise<AxiosResponse<TxResponse>> => {
  try {
    const response = await axiosInstance.post("/tx", req);
    return response;
  } catch (error) {
    console.error(error);
    return {} as AxiosResponse<TxResponse>;
  }
};