import { Count, KVCount, KeyValuePairString } from "@/types/buckets";
import { BucketsApi } from './interface'
import { http } from '../http'

const Buckets_Url = '/buckets'

export const createBucketsApi = (): BucketsApi => ({
  changeDataSource: async (file: string) => await http.post<void>(`${Buckets_Url}/data/change`, { path: file }),
  listDataSource: async () => await http.get<string[]>(`${Buckets_Url}/data/list`),
  getBuckets: async () => await http.get<string[]>(Buckets_Url),
  getKeysCount: async (bucket: string) => await http.get<Count>(`${Buckets_Url}/${bucket}/count`),
  getKeysCountLength: async (bucket: string, length: number) => await http.get<KVCount>(`${Buckets_Url}/${bucket}/count/${length}/keys`),
  getValuesCountLength: async (bucket: string, length: number) => await http.get<KVCount>(`${Buckets_Url}/${bucket}/count/${length}/values`),
  getPages: async (bucket: string, pageNumber: number, resultsNumber: number) => await http.get<KeyValuePairString[]>(`${Buckets_Url}/${bucket}/pages/${pageNumber}/${resultsNumber}`),
  searchByKey: async (bucket: string, key: string) => await http.get<KeyValuePairString[]>(`${Buckets_Url}/${bucket}/keys/${key}`),
  searchByValue: async (bucket: string, value: string) => await http.get<KeyValuePairString[]>(`${Buckets_Url}/${bucket}/values/${value}`)
})