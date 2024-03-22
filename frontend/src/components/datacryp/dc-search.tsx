"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { useApi } from "@/api/api";
import { useBucketContext } from "@/context/bucket-context";
import { Button } from "@/components/ui/button";
import { SearchIcon } from "lucide-react";
import { useRouter } from "next/navigation";

export default function DatacrypSearch() {
  const router = useRouter();
  const api = useApi();
  const {
    selectedBucket,
  } = useBucketContext();
  const [inputKeySearchValue, setInputKeySearchValue] = useState<string>("");
  const [inputValueSearchValue, setInputValueSearchValue] =
    useState<string>("");

  const searchKeysByLength = async (key: string) => {
    if (!isNaN(Number(key))) {
      router.push(`/dataview/key?bucketName=${selectedBucket}&key=${key}`)
    } else {
      console.error('Key is not a number');
    }
  };

  const searchValuesByLength = async (value: string) => {
    if (!isNaN(Number(value))) {
      router.push(`/dataview/value?bucketName=${selectedBucket}&value=${value}`)
    } else {
      console.error('Value is not a number');
    }
  };

  return (
    <div className="flex items-center justify-between w-full gap-4">
      <div className="w-full p-1 border border-input rounded-full">
        <div className="w-full relative">
          <Input
            disabled={selectedBucket == ""}
            className="w-full rounded-full h-10 px-4 border-0"
            placeholder="Search by key length..."
            value={inputKeySearchValue}
            onChange={(e) => setInputKeySearchValue(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                searchKeysByLength(inputKeySearchValue);
              }
            }}
          />
          <Button
            disabled={selectedBucket == ""}
            className="absolute right-0 top-0 rounded-full h-10 w-10"
            variant="ghost"
            onClick={() => searchKeysByLength(inputKeySearchValue)}
          >
            <SearchIcon className="absolute w-4" />
          </Button>
        </div>
      </div>
      <div className="w-full flex items-center justify-center">
        <h1 className="font-bold text-2xl">Dataview <span className="font-thin">MDBX viewer</span></h1>
      </div>
      <div className="w-full p-1 border border-input rounded-full">
        <div className="w-full relative">
          <Input
            disabled={selectedBucket == ""}
            className="w-full rounded-full h-10 px-4 border-0"
            placeholder="Search by value length..."
            value={inputValueSearchValue}
            onChange={(e) => setInputValueSearchValue(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                searchValuesByLength(inputValueSearchValue);
              }
            }}
          />
          <Button
            disabled={selectedBucket == ""}
            className="absolute right-0 top-0 rounded-full h-10 w-10"
            variant="ghost"
            onClick={() => searchValuesByLength(inputValueSearchValue)}
          >
            <SearchIcon className="absolute w-4" />
          </Button>
        </div>
      </div>
    </div>
  );
}
