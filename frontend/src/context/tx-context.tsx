'use client'

import React, {
  createContext,
  useContext,
  useState,
} from 'react'
import { TxOutput } from '@/types/tx'

interface TxContextProps {
  response: TxOutput | string;
  setResponse: React.Dispatch<React.SetStateAction<TxOutput | string>>;
}

const TxContext = createContext<TxContextProps>({} as TxContextProps)

export const TxProvider = ({ children }: { children: React.ReactNode }) => {
  const [response, setResponse] = useState<TxOutput | string>("Send a transaction to see the response.")

  const contextValue: TxContextProps = {
    response,
    setResponse,
  }


  return (
    <TxContext.Provider value={contextValue}>{children}</TxContext.Provider>
  )
};

export const useTxContext = () => useContext(TxContext)