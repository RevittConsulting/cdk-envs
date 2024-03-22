'use client'

import React, {
  createContext,
  useContext,
  useState,
} from 'react'
import { Count, KeyValuePairString, KVCount } from "@/types/buckets";

interface BucketContextProps {
  loading: boolean;
  setLoading: React.Dispatch<React.SetStateAction<boolean>>;
  selectedFile: string;
  setSelectedFile: React.Dispatch<React.SetStateAction<string>>;
  selectedBucket: string;
  setSelectedBucket: React.Dispatch<React.SetStateAction<string>>;
  pages: KeyValuePairString[];
  setPages: React.Dispatch<React.SetStateAction<KeyValuePairString[]>>;
  buckets: string[];
  setBuckets: React.Dispatch<React.SetStateAction<string[]>>;
  count: Count;
  setCount: React.Dispatch<React.SetStateAction<Count>>;
  resultsInput: string;
  setResultsInput: React.Dispatch<React.SetStateAction<string>>;
  pageNumInput: string;
  setPageNumInput: React.Dispatch<React.SetStateAction<string>>;
  kvs: KVCount | null;
  setKvs: React.Dispatch<React.SetStateAction<KVCount | null>>;
}

const BucketContext = createContext<BucketContextProps | undefined>(undefined)

export const BucketProvider = ({ children }: { children: React.ReactNode }) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [selectedFile, setSelectedFile] = useState<string>("");
  const [selectedBucket, setSelectedBucket] = useState<string>("");
  const [pages, setPages] = useState<KeyValuePairString[]>([]);
  const [kvs, setKvs] = useState<KVCount | null>(null);
  const [buckets, setBuckets] = useState<string[]>([]);
  const [count, setCount] = useState<Count>({ count: 0 });
  const [resultsInput, setResultsInput] = useState<string>("100");
  const [pageNumInput, setPageNumInput] = useState<string>("0");

  return (
    <BucketContext.Provider value={
      {
        loading,
        setLoading,
        selectedFile,
        setSelectedFile,
        selectedBucket,
        setSelectedBucket,
        pages,
        setPages,
        buckets,
        setBuckets,
        count,
        setCount,
        resultsInput,
        setResultsInput,
        pageNumInput,
        setPageNumInput,
        kvs,
        setKvs,
      }
    }>{children}</BucketContext.Provider>
  )
};

export function useBucketContext() {
  const context = useContext(BucketContext);
  if (context === undefined) {
    throw new Error('useBuckets must be used within a BucketsProvider');
  }
  return context;
}