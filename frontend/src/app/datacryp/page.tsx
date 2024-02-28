'use client';

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import axios from "axios";

export default function DatacrypPage() {
  const [buckets, setBuckets] = useState([]);

  const getBuckets = async () => {
    const response = await axios.get("http://localhost:8080/api/v1/buckets");
    const data = response.data;
    console.log(data);
    setBuckets(data);
  }


  return (
    <div className="flex flex-col items-center justify-center gap-4 flex-wrap">
      <h1 className="my-10">Buckets</h1>
      <div className="flex flex-row gap-4">
        <Button onClick={getBuckets}>Load Buckets</Button>
        <Select>
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Buckets" />
          </SelectTrigger>
          <SelectContent>
            {buckets.map((bucket) => (
              <SelectItem key={bucket} value={bucket}>{bucket}</SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>
    </div>
  );
}
