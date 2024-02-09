'use client'

import { useState, useEffect } from 'react';
import axios from 'axios';
import { Block } from '../types/block';

export default function BlockCard() {

  const [blockNumber, setBlockNumber] = useState<number | null>(null);

  const getHighestBlock = async () => {
    try {
      const response = await axios.get<Block>('http://localhost:80/api/v1/chain/block');
      const data = response.data;
      setBlockNumber(data.number);
    } catch (error) {
      console.error('Error fetching block data:', error);
    }
  };

  useEffect(() => {
    getHighestBlock();
  }, []);


  return (
    <div className="max-w-sm rounded overflow-hidden shadow-sm p-4 m-4 bg-white border border-gray-200">
      <div className="font-bold text-xl mb-2">Block Information</div>
      <div className="text-gray-700 text-base">
        {blockNumber !== null ? `Highest Block Number: ${blockNumber}` : 'Loading...'}
      </div>
      <button
        className="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onClick={getHighestBlock}
      >
        Refresh Data
      </button>
    </div>
  )
}
