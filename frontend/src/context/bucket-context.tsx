'use client'

import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
} from 'react'
import { KeyValuePairString } from "@/types/buckets";

interface BucketContextProps {
  selectedFile: string;
  setSelectedFile: (file: string) => void;
  selectedBucket: string;
  setSelectedBucket: (bucket: string) => void;
  pages: KeyValuePairString[];
  setPages: (pages: KeyValuePairString[]) => void;
}

const BucketContext = createContext<BucketContextProps>({} as BucketContextProps)

export const BucketProvider = ({ children }: { children: React.ReactNode }) => {
  const [selectedFile, setSelectedFile] = useState<string>("");
  const [selectedBucket, setSelectedBucket] = useState<string>("");
  const [pages, setPages] = useState<KeyValuePairString[]>([]);

  const contextValue: BucketContextProps = {
    selectedFile,
    setSelectedFile,
    selectedBucket,
    setSelectedBucket,
    pages,
    setPages,
  }

  return (
    <BucketContext.Provider value={contextValue}>{children}</BucketContext.Provider>
  )
};

export const useBucketContext = () => useContext(BucketContext)