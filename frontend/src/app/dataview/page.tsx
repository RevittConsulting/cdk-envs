"use client";

import DatacrypHeader from "@/components/datacryp/dc-header";
import DatacrypSearch from "@/components/datacryp/dc-search";
import DataTable from "@/components/datacryp/data-table";

export default function DatacrypPage() {
  return (
    <div className="flex flex-col items-center justify-center gap-4 flex-wrap w-full px-4 py-4">
      {/* <DatacrypSearch /> */}
      <DatacrypHeader />
      <DataTable />
    </div>
  );
}
