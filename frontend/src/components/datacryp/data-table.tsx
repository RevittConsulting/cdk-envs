"use client";

import { useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { useBucketContext } from "@/context/bucket-context";
import { Button } from "@/components/ui/button";
import { CornerRightUp } from "lucide-react";
import { useApi } from "@/api/api";
import { useRouter } from "next/navigation";
import { RefreshCcw } from "lucide-react";
import { Undo2 } from "lucide-react";

export default function DataTable() {
  const router = useRouter();
  const api = useApi();
  const { pages, setPages, selectedBucket, loading, setLoading, pageNumInput, resultsInput } =
    useBucketContext();

  const searchKeys = async (key: string) => {
    setPages([]);
    setLoading(true);
    const res = await api.buckets.searchByKey(selectedBucket, key);
    if (res.data) {
      setPages(res.data);
    }
    setLoading(false);
  };

  const searchValues = async (value: string) => {
    setPages([]);
    setLoading(true);
    const res = await api.buckets.searchByValue(selectedBucket, value);
    if (res.data) {
      setPages(res.data);
    }
    setLoading(false);
  };

  const searchKeysByLength = async (key: number) => {
    if (!isNaN(Number(key))) {
      router.push(`/dataview/key?bucketName=${selectedBucket}&key=${key}`)
    } else {
      console.error('Key is not a number');
    }
  };

  const searchValuesByLength = async (value: number) => {
    if (!isNaN(Number(value))) {
      router.push(`/dataview/value?bucketName=${selectedBucket}&value=${value}`)
    } else {
      console.error('Value is not a number');
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

  if (loading) {
    return (
      <div className="rounded-md border w-full">
        <Table>
          <TableHeader className="bg-gray-500/10">
            <TableRow>
              <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-1/24 border-r">
              </TableHead>
              <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-11/24 border-r">
                <p>Key</p>
              </TableHead>
              <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-12/24">
                <p>Value</p>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell colSpan={3} className="text-center py-4 text-gray-500">
                Loading...
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>
    );
  }

  return (
    <div className="rounded-md border w-full h-full">
      <Table className="h-full">
        <TableHeader className="bg-gray-500/10">
          <TableRow>
            <TableHead className="py-1 px-1 h-auto text-left text-sm font-semibold w-1/24 border-r">
            <Button
            variant="ghost"
            onClick={getResults}
            disabled={selectedBucket == ""}
          >
            {selectedBucket != "" &&
              <Undo2 className="w-4" />}
          </Button>
            </TableHead>
            <TableHead className="py-1 px-3 h-auto text-left text-sm font-semibold w-11/24 border-r">
              <p>Key</p>
            </TableHead>
            <TableHead className="py-1 px-3 h-auto text-left text-sm font-semibold w-12/24">
              <p>Value</p>
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody className="h-full">
          {pages.length > 0 ? (
            pages.map((page, index) => (
              <TableRow key={index} className="h-full">
                <TableCell className="px-3 py-4 text-sm border-r">
                  {index + 1}
                </TableCell>
                <TableCell className="px-3 py-0 text-sm border-r h-full">
                  <div className="flex justify-between items-start h-full gap-2">

                    <div className="h-full flex items-center justify-center">
                    <button onClick={() => searchKeys(page.key)}>
                      <Tooltip delayDuration={100}>
                        <TooltipTrigger>
                          <p style={{ wordBreak: "break-word" }}>{page.key}</p>
                        </TooltipTrigger>
                        <TooltipContent side="top">
                          <p>{`length: ${page.key.length}`}</p>
                        </TooltipContent>
                      </Tooltip>
                    </button>
                    </div>

                    <Tooltip delayDuration={100}>
                      <TooltipTrigger>
                        <Button
                          className="py-1 px-2.5 selection: h-auto my-2"
                          variant="ghost"
                          onClick={() =>
                            searchKeysByLength(page.key.length)
                          }
                        >
                          <CornerRightUp className="w-3" />
                        </Button>
                      </TooltipTrigger>
                      <TooltipContent side="top">
                        <p>Search by key length</p>
                      </TooltipContent>
                    </Tooltip>
                  </div>
                </TableCell>
                <TableCell className="px-3 py-0 text-sm h-full">
                  <div className="flex justify-between items-start h-full gap-2">

                  <div className="h-full flex items-center justify-center">
                    <button onClick={() => searchValues(page.value)}>
                      <Tooltip delayDuration={100}>
                        <TooltipTrigger>
                          <p style={{ wordBreak: "break-word" }} className="my-2">
                            {page.value}
                          </p>
                        </TooltipTrigger>
                        <TooltipContent side="top">
                          <p>{`length: ${page.value.length}`}</p>
                        </TooltipContent>
                      </Tooltip>
                    </button>
                    </div>
                    <Tooltip delayDuration={100}>
                      <TooltipTrigger>
                        <Button
                          className="py-1 px-2.5 selection: h-auto my-2"
                          variant="ghost"
                          onClick={() =>
                            searchValuesByLength(page.value.length)
                          }
                        >
                          <CornerRightUp className="w-3" />
                        </Button>
                      </TooltipTrigger>
                      <TooltipContent side="top">
                        <p>Search by value length</p>
                      </TooltipContent>
                    </Tooltip>
                  </div>
                </TableCell>
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={3} className="text-center py-4 text-gray-500">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
