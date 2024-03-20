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
import { Copy } from "lucide-react";
import { useApi } from "@/api/api";

export default function DataTable() {
  const api = useApi();
  const { pages, setPages, selectedBucket, loading, setLoading } =
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

  if (loading) {
    return (
      <div className="rounded-md border w-full">
        <Table>
          <TableHeader className="bg-gray-500/10">
            <TableRow>
              <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-1/24 border-r"></TableHead>
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
    <div className="rounded-md border w-full">
      <Table>
        <TableHeader className="bg-gray-500/10">
          <TableRow>
            <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-1/24 border-r"></TableHead>
            <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-11/24 border-r">
              <p>Key</p>
            </TableHead>
            <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-12/24">
              <p>Value</p>
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {pages.length > 0 ? (
            pages.map((page, index) => (
              <TableRow key={index}>
                <TableCell className="px-3 py-4 text-sm border-r">
                  {index + 1}
                </TableCell>
                <TableCell className="px-3 py-0 text-sm border-r">
                  <div className="flex justify-between items-center">
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
                    <Button
                      className="py-1 px-2.5 selection: h-auto"
                      variant="ghost"
                      onClick={() => navigator.clipboard.writeText(page.key)}
                    >
                      <Copy className="w-3" />
                    </Button>
                  </div>
                </TableCell>
                <TableCell className="px-3 py-0 text-sm">
                  <div className="flex justify-between items-center">
                    <button onClick={() => searchValues(page.value)}>
                      <Tooltip delayDuration={100}>
                        <TooltipTrigger>
                          <p style={{ wordBreak: "break-word" }}>
                            {page.value}
                          </p>
                        </TooltipTrigger>
                        <TooltipContent side="top">
                          <p>{`length: ${page.value.length}`}</p>
                        </TooltipContent>
                      </Tooltip>
                    </button>
                    <Button
                      className="py-1 px-2.5 selection: h-auto"
                      variant="ghost"
                      onClick={() => navigator.clipboard.writeText(page.value)}
                    >
                      <Copy className="w-3" />
                    </Button>
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
