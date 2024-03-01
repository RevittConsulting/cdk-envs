"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { Count, KeyValuePairString } from "@/types/buckets";
import { getBuckets, getCount, loadPages, loadKeys } from "@/api/buckets";
import { useBucketContext } from "@/context/bucket-context";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label"
import { ScrollArea } from "@/components/ui/scroll-area";
import DatacrypHeader from "@/components/datacryp/dc-header";
import DataTable from "@/components/datacryp/data-table";

export default function DatacrypPage() {
  const { selectedBucket } = useBucketContext();

  const [resultsInput, setResultsInput] = useState<string>("100");
  const [pageNumInput, setPageNumInput] = useState<string>("1");
  const [keyInput, setKeyInput] = useState<string>("");



  const getKeys = async () => {
    const data = await loadKeys(selectedBucket, keyInput);
  };

  return (
    <div className="flex flex-col items-center justify-center gap-4 flex-wrap w-full px-4 pt-4">
      
      <DatacrypHeader />

      <DataTable />

      {/* <div className="flex flex-row gap-4 items-center w-full">
        <div className="flex items-center space-x-2">
          <Label htmlFor="pageNum" className="whitespace-nowrap">Page Number</Label>
          <Input type="number" id="pageNum" value={pageNumInput} min="1" onChange={(e) => setPageNumInput(e.target.value)} />
        </div>
        <div className="flex items-center space-x-2">
          <Label htmlFor="results" className="whitespace-nowrap">Results per Page</Label>
          <Input type="number" id="results" value={resultsInput} min="1" onChange={(e) => setResultsInput(e.target.value)} />
        </div>
        <Button onClick={getPages}>Load Pages</Button>
      </div>

      <div className="flex flex-row gap-4 items-center w-full">
        <ScrollArea className="p-4 whitespace-nowrap">
          <div className="whitespace-pre w-max">{JSON.stringify(pages, null, 2)}</div>
        </ScrollArea>
      </div>

      <div className="flex flex-row gap-4 items-center w-full">
        <div className="flex items-center space-x-2">
          <Label htmlFor="keyInput" className="whitespace-nowrap">Key Value</Label>
          <Input type="text" id="keyInput" value={keyInput} onChange={(e) => setKeyInput(e.target.value)} />
        </div>
        <Button onClick={getKeys}>Load Pages</Button>
      </div> */}
    </div>
  );
}
