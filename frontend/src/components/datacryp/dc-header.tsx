"use client";

import { useState } from "react";
import { Count } from "@/types/buckets";
import { getBuckets, getCount, loadPages } from "@/api/buckets";
import { useBucketContext } from "@/context/bucket-context";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { FileIcon } from "@radix-ui/react-icons";

export default function DatacrypHeader() {
  const { selectedFile, selectedBucket, setSelectedBucket, setPages } =
    useBucketContext();
  const [buckets, setBuckets] = useState<string[]>([]);
  const [count, setCount] = useState<Count>({ count: 0 });

  const fetchBuckets = async () => {
    const data = await getBuckets();
    setBuckets(data);
  };

  const setBucket = async (bucket: string) => {
    setSelectedBucket(bucket);

    const countRes = await getCount(bucket);
    setCount(countRes);

    const pagesRes = await loadPages(bucket, 1, 100);
    setPages(pagesRes);
  }

  return (
    <div className="flex sm:flex-row flex-col items-center justify-between gap-4 w-full">
      <div className="w-full">
        <div className="flex flex-row gap-4 justify-center md:justify-start">
          {selectedFile ? (
            <div className="border p-2 px-3 rounded-md">
              <h1 className="flex items-center gap-1">
                <span className="mr-2">Data:</span>
                <FileIcon className="h-4 w-4" aria-hidden="true" />{" "}
                {selectedFile.split("/").pop()}
              </h1>
            </div>
          ) : (
            <div className="border p-2 px-3 rounded-md">
              <h1 className="text-gray-500">Select a data source.</h1>
            </div>
          )}
        </div>
      </div>

      <div className="w-full">
        <div className="flex flex-row gap-4 justify-center">
          <Button disabled={selectedFile == ""} onClick={fetchBuckets}>
            Load Buckets
          </Button>
          <Select
            onValueChange={(value) => setBucket(value)}
            disabled={buckets.length < 1}
          >
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Buckets" />
            </SelectTrigger>
            <SelectContent>
              {buckets.map((bucket) => (
                <SelectItem key={bucket} value={bucket}>
                  {bucket}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>
      </div>

      <div className="w-full">
        <div className="flex justify-center md:justify-end gap-4 items-center">
          <div className="border p-2 px-3 rounded-md">
            {selectedFile ? (
              <span>Key value pairs: {count.count}</span>
            ) : (
              <span className="text-gray-500">No data source selected.</span>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
