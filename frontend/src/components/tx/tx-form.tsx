"use client";

import { useState, FormEvent } from "react";
import { useApi } from "@/api/api";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import { Send } from "lucide-react";
import { useTxContext } from "@/context/tx-context";

export interface TxFormData {
  host: string;
  toAddress: string;
  fromAddress: string;
  privateKey: string;
  amount: number;
  gasLimit: number;
  gasPrice: number;
}

export default function TxForm() {
  const api = useApi();
  const { setResponse } = useTxContext();
  const [formData, setFormData] = useState<TxFormData>({
    host: "http://localhost:8545",
    toAddress: "",
    fromAddress: "",
    privateKey: "",
    amount: 10000,
    gasLimit: 21000,
    gasPrice: 5000,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    if (name === 'amount' || name === 'gasLimit' || name === 'gasPrice') {
      setFormData(prevState => ({ ...prevState, [name]: Number(value) }));
    } else {
      setFormData(prevState => ({ ...prevState, [name]: value }));
    }
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    const res = await api.tx.submitTx(formData);
    if (res.data?.output) {
      setResponse(res.data.output);
    } else if (res.data?.error) {
      setResponse(res.data.error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="w-[30vw] flex flex-col gap-4">
      <div>
        <Label htmlFor="host">Host:</Label>
        <Input
          type="text"
          name="host"
          value={formData.host}
          onChange={handleChange}
        />
      </div>
      <div>
        <Label htmlFor="toAddress">To Address:</Label>
        <Input
          type="text"
          name="toAddress"
          value={formData.toAddress}
          onChange={handleChange}
        />
      </div>
      <div>
        <Label htmlFor="toAddress">From Address:</Label>
        <Input
          type="text"
          name="fromAddress"
          value={formData.fromAddress}
          onChange={handleChange}
        />
      </div>
      <div>
        <Label htmlFor="toAddress">Private Key:</Label>
        <Input
          type="text"
          name="privateKey"
          value={formData.privateKey}
          onChange={handleChange}
        />
      </div>
      <div className="flex gap-4">
        <div>
          <Label htmlFor="amount">Amount:</Label>
          <Input
            type="number"
            name="amount"
            value={formData.amount}
            onChange={handleChange}
          />
        </div>
        <div>
          <Label htmlFor="gasLimit">Gas Limit:</Label>
          <Input
            type="number"
            name="gasLimit"
            value={formData.gasLimit}
            onChange={handleChange}
          />
        </div>
        <div>
          <Label htmlFor="gasPrice">Gas Price:</Label>
          <Input
            type="number"
            name="gasPrice"
            value={formData.gasPrice}
            onChange={handleChange}
          />
        </div>
      </div>
      <Button type="submit" className="w-full">
        Send Tx <Send size={12} className="ml-2" />
      </Button>
    </form>
  );
}
