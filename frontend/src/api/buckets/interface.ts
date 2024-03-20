import { ApiResponse } from '../types'
import { Count, KeyValuePairString } from "@/types/buckets";

export interface BucketsApi {
  changeDataSource: (file: string) => Promise<ApiResponse<void>>
  listDataSource: () => Promise<ApiResponse<string[]>>
  getBuckets: () => Promise<ApiResponse<string[]>>
  getKeysCount: (bucket: string) => Promise<ApiResponse<Count>>
  getPages: (bucket: string, pageNumber: number, resultsNumber: number) => Promise<ApiResponse<KeyValuePairString[]>>
  searchByKey: (bucket: string, key: string) => Promise<ApiResponse<KeyValuePairString[]>>
  searchByValue: (bucket: string, value: string) => Promise<ApiResponse<KeyValuePairString[]>>
}