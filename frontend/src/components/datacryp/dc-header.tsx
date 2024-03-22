"use client";

import { useApi } from "@/api/api";
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
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function DatacrypHeader() {
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

  const setBucket = async (bucket: string) => {
    setSelectedBucket(bucket);
    setPages([]);
    setCount({ count: 0 });

    const countRes = await api.buckets.getKeysCount(bucket);
    if (countRes.data) {
      setCount(countRes.data);
    }

    const pagesRes = await api.buckets.getPages(
      bucket,
      Number(pageNumInput),
      Number(resultsInput)
    );
    if (pagesRes.data) {
      setPages(pagesRes.data);
    }
  };

  const getResults = async () => {
    const pagesRes = await api.buckets.getPages(
      selectedBucket,
      Number(pageNumInput),
      Number(resultsInput)
    );
    if (pagesRes.data) {
      setPages(pagesRes.data);
    } else {
      setPages([]);
    }
  };

  const handlePageNumChange = async (pageNum: string) => {
    setPageNumInput(pageNum);

    const pagesRes = await api.buckets.getPages(
      selectedBucket,
      Number(pageNum),
      Number(resultsInput)
    );
    if (pagesRes.data) {
      setPages(pagesRes.data);
    }
  };

  const handleResultsChange = async (results: string) => {
    setResultsInput(results);

    const pagesRes = await api.buckets.getPages(
      selectedBucket,
      Number(pageNumInput),
      Number(results)
    );
    if (pagesRes.data) {
      setPages(pagesRes.data);
    }
  };

  return (
    <div className="flex sm:flex-row flex-col items-center justify-between gap-4 w-full">
      <div className="w-full">
        <div className="flex flex-row gap-4 justify-center md:justify-start">
          {selectedFile ? (
            <div className="border p-2 px-3 rounded-md">
              <h1 className="flex items-center gap-1">
                <span className="mr-2">Data:</span>
                <FileIcon className="h-4 w-4" aria-hidden="true" />{" "}
                {selectedFile.split("/").slice(-2).join("/")}
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
          <div className="flex items-center justify-center gap-2">
            <Label className={selectedBucket == "" ? "opacity-50" : ""}>
              Page
            </Label>
            <Input
              type="number"
              disabled={selectedBucket == ""}
              value={pageNumInput}
              onChange={(e) => handlePageNumChange(e.target.value)}
              className="w-20"
              min={0}
              max={Math.ceil(count.count / Number(resultsInput))}
            />
          </div>
          <div className="flex items-center justify-center gap-2">
            <Label className={selectedBucket == "" ? "opacity-50" : ""}>
              Results
            </Label>
            <Input
              type="number"
              disabled={selectedBucket == ""}
              value={resultsInput}
              onChange={(e) => {
                const value = Number(e.target.value);
                if (value <= 200) {
                  handleResultsChange(e.target.value);
                }
              }}
              className="w-20"
              min={1}
              max={200}
            />
          </div>
          {/* <Button
            variant="outline"
            onClick={getResults}
            disabled={selectedBucket == ""}
          >
            
          </Button> */}
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
