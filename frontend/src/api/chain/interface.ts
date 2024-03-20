import { ApiResponse } from "../types";
import { Chain } from "@/types/chain";

export interface ChainApi {
  getChains: () => Promise<ApiResponse<Chain[]>>;
  restartServices: (serviceName: string) => Promise<ApiResponse<void>>;
  stopAllServices: () => Promise<ApiResponse<void>>;
}