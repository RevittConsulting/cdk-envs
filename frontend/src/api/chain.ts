import { Chain } from "@/types/chain";
import { axiosInstance } from '@/utils/axios-instance'

export const getChains = async (): Promise<Chain[]> => {
  try {
    const response = await axiosInstance.get<Chain[]>("/chains");
    if (response.status === 200) {
      return response.data;
    }
  } catch (error) {
    console.error(error);
    return [];
  }
  return [];
};

export const restartServices = async (serviceName: string) => {
  try {
    const response = await axiosInstance.post("/chains", { serviceName });
    if (response.status === 200) {
      return response.data;
    }
  } catch (error) {
    console.error(error);
    return {};
  }
  return {};
};