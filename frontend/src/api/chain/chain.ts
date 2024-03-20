import { Chain } from "@/types/chain";
import { ChainApi } from './interface'
import { http } from '../http'

const Chain_Url = '/chain'

export const createChainApi = (): ChainApi => ({
  getChains: async () => await http.get<Chain[]>("/chains"),
  restartServices: async (serviceName: string) => await http.post<void>("/chains", { serviceName }),
  stopAllServices: async () => await http.get<void>("/chains/stop")
})