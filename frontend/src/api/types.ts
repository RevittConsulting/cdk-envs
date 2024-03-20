import { BucketsApi } from './buckets/interface'
import { ChainApi } from './chain/interface'
import { DatastreamApi } from './datastream/interface'
import { TxApi } from './tx/interface'

export type ApiResponse<T> = {
	data: T | null
	status?: number
	error: T | null | Error
}

export interface ApiFactory {
	buckets: BucketsApi
	chain: ChainApi
  datastream: DatastreamApi
  tx: TxApi
}