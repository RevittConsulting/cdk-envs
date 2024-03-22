"use client";

import { useEffect } from "react";
import KVTable from "@/components/datacryp/kv-table";
import { Button } from "@/components/ui/button";
import { useRouter, useSearchParams } from "next/navigation";
import { Undo2 } from "lucide-react";
import { useBucketContext } from "@/context/bucket-context";
import { useApi } from "@/api/api";

export default function ValuePage() {
  const api = useApi();
  const router = useRouter();
  const searchParams = useSearchParams();
  const value = searchParams.get("value");
  const {
    selectedBucket,
    setLoading,
    kvs,
    setKvs,
  } = useBucketContext();

  useEffect(() => {
    console.log("value: ", value);
    getKVs();
  }, []);

  const getKVs = async () => {
    setLoading(true);
    const res = await api.buckets.getValuesCountLength(selectedBucket, Number(value));
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
          <h1>Total Values At Length {value} <span className="font-thin">- limit to 1000</span></h1>
        </div>
        <div className="w-full">
          <div className="flex justify-center md:justify-end gap-4 items-center">
            <div className="border p-2 px-3 rounded-md">
              {kvs ? (
                <span>Values count: {kvs.count}</span>
              ) : (
                <span className="text-gray-500">No Values.</span>
              )}
            </div>
          </div>
        </div>
      </div>
      <KVTable />
    </div>
  );
}