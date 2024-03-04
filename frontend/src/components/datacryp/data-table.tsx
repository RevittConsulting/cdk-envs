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
import { KeyValuePairString } from "@/types/buckets";
import { useBucketContext } from "@/context/bucket-context";

export default function DataTable() {
  const { pages } = useBucketContext();

  return (
    <div className="rounded-md border w-full">
      <Table>
        <TableHeader className="bg-gray-500/10">
          <TableRow>
            <TableHead
              className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold sm:pl-6 w-1/2 border-r"
            >
              Key
            </TableHead>
            <TableHead
              className="px-3 py-3.5 text-left text-sm font-semibold w-2/2"
            >
              Value
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {pages.length > 0 ? (
            pages.map((page, index) => (
              <TableRow key={index}>
                <TableCell className="px-3 py-4 text-sm border-r">
                  {page.key}
                </TableCell>
                <TableCell className="px-3 py-4 text-sm">
                  {page.value}
                </TableCell>
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={2} className="text-center py-4 text-gray-500">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
