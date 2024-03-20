'use client';

import { useState } from 'react';
import { Input } from '@/components/ui/input'
import { useApi } from '@/api/api'
import { useBucketContext } from '@/context/bucket-context'

export default function DatacrypSearch() {
  const api = useApi();
  const {
    selectedFile,
    selectedBucket,
    setSelectedBucket,
    setPages,
    buckets,
    setBuckets,
    count,
    setCount,
    resultsInput,
    setResultsInput,
    pageNumInput,
    setPageNumInput,
  } = useBucketContext();
  const [inputValue, setInputValue] = useState<string>('');

  const searchKeys = async (key: string) => {
    const res = await api.buckets.searchByKey(selectedBucket, key);
    if (res.data) {
      setPages(res.data);
    }
  };

  return (
    <div>
      <Input 
        className='w-80' 
        value={inputValue}
        onChange={(e) => setInputValue(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === 'Enter') {
            searchKeys(inputValue);
          }
        }}
      />
    </div>
  )
}
