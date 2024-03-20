import { ApiFactory } from './types';
import { createBucketsApi } from './buckets/buckets';
import { createChainApi } from './chain/chain';
import { createDatastreamApi } from './datastream/datastream';
import { createTxApi } from './tx/tx';

export const useApi = (): ApiFactory => {
  return {
    buckets: createBucketsApi(),
    chain: createChainApi(),
    datastream: createDatastreamApi(),
    tx: createTxApi(),
  }
};