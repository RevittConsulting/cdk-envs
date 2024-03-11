'use client'

import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
} from 'react'

interface ChainContextProps {
  activeChain: string | null
  setActiveChain: (chain: string | null) => void
}

const ChainContext = createContext<ChainContextProps>({} as ChainContextProps)

export const ChainProvider = ({ children }: { children: React.ReactNode }) => {
  const [activeChain, setActiveChain] = useState<string | null>(null)

  const contextValue: ChainContextProps = {
    activeChain,
    setActiveChain,
  }

  return (
    <ChainContext.Provider value={contextValue}>{children}</ChainContext.Provider>
  )
};

export const useChainContext = () => useContext(ChainContext)