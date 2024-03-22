import { ApiResponse } from '../types'
import { Count, KVCount, KeyValuePairString } from "@/types/buckets";

export interface BucketsApi {
  changeDataSource: (file: string) => Promise<ApiResponse<void>>
  listDataSource: () => Promise<ApiResponse<string[]>>
  getBuckets: () => Promise<ApiResponse<string[]>>
  getKeysCount: (bucket: string) => Promise<ApiResponse<Count>>
  getKeysCountLength: (bucket: string, length: number) => Promise<ApiResponse<KVCount>>
  getValuesCountLength: (bucket: string, length: number) => Promise<ApiResponse<KVCount>>
  getPages: (bucket: string, pageNumber: number, resultsNumber: number) => Promise<ApiResponse<KeyValuePairString[]>>
  searchByKey: (bucket: string, key: string) => Promise<ApiResponse<KeyValuePairString[]>>
  searchByValue: (bucket: string, value: string) => Promise<ApiResponse<KeyValuePairString[]>>
}