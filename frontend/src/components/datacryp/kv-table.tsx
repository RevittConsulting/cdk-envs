"use client";

import React from 'react'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useApi } from "@/api/api";
import { useBucketContext } from "@/context/bucket-context";

export default function KVTable() {
  const api = useApi();
  const { selectedBucket, loading, setLoading, kvs, setKvs } =
    useBucketContext();

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
            <TableHead className="py-3 px-4 h-auto text-left text-sm font-semibold w-23/24 border-r">Values</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {kvs && kvs.kv && kvs.kv.length > 0 ? (
            kvs.kv.map((kv, index) => (
              <TableRow key={index}>
                <TableCell className="px-3 py-2 text-sm border-r">
                  {index + 1}
                </TableCell>
                <TableCell className="px-3 py-2 text-sm">
                  <div className=''>{kv}</div>
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
  )
}
