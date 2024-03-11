import { Chain } from "@/types/chain";
import { axiosInstance } from '@/utils/axios-instance'
import { AxiosResponse } from 'axios';

export const getChains = async (): Promise<AxiosResponse<Chain[]>> => {
  try {
    const response = await axiosInstance.get<Chain[]>("/chains");
    return response
  } catch (error) {
    console.error(error);
    return {} as AxiosResponse<Chain[]>;
  }
};

export const restartServices = async (serviceName: string): Promise<AxiosResponse> => {
  try {
    const response = await axiosInstance.post("/chains", { serviceName });
    return response;
  } catch (error) {
    console.error(error);
    return {} as AxiosResponse;
  }
};

export const stopAllServices = async (): Promise<AxiosResponse> => {
  try {
    const response = await axiosInstance.get("/chains/stop");
    return response;
  } catch (error) {
    console.error(error);
    return {} as AxiosResponse;
  }
}