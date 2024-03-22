"use client";

import React, { useEffect } from "react";
import KVTable from "@/components/datacryp/kv-table";
import { Button } from "@/components/ui/button";
import { useSearchParams } from 'next/navigation';
import { useRouter } from 'next/navigation';
import { Undo2 } from "lucide-react";
import { useBucketContext } from "@/context/bucket-context";
import { useApi } from "@/api/api";

export default function KeyPage() {
  const api = useApi();
  const router = useRouter();
  const searchParams = useSearchParams();
  const key = searchParams.get("key");
  const {
    selectedBucket,
    setLoading,
    kvs,
    setKvs,
  } = useBucketContext();

  useEffect(() => {
    console.log("key: ", key);
    getKVs();
  }, []);

  const getKVs = async () => {
    setLoading(true);
    const res = await api.buckets.getKeysCountLength(selectedBucket, Number(key));
    if (res.data) {
      setKvs(res.data);
    }
    setLoading(false);
  }

  const goBack = () => {
    router.push("/dataview");
  };

  return (
    <div className="w-full flex flex-col items-center justify-center gap-4 px-4 py-4">
      <div className="w-full flex items-center justify-between">
        <div className="w-full">
          <Button variant="outline" onClick={goBack}>
            <Undo2 className="w-4 mr-2" /> Go back
          </Button>
        </div>
        <div className="w-full flex items-center justify-center">
          <h1>Total Keys At Length {key} <span className="font-thin">- limit to 1000</span></h1>
        </div>
        <div className="w-full">
          <div className="flex justify-center md:justify-end gap-4 items-center">
            <div className="border p-2 px-3 rounded-md">
              {kvs ? (
                <span>Keys count: {kvs.count}</span>
              ) : (
                <span className="text-gray-500">No Keys.</span>
              )}
            </div>
          </div>
        </div>
      </div>
      <KVTable />
    </div>
  );
}