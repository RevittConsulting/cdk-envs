"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { FileIcon } from "@radix-ui/react-icons";
import { useBucketContext } from "@/context/bucket-context";
import { useApi } from "@/api/api";
import { Database } from "lucide-react";

export default function Sidebar() {
  const api = useApi();
  const [files, setFiles] = useState<string[]>([]);
  const { selectedFile, setSelectedFile } = useBucketContext();

  const fetchDbFiles = async () => {
    const res = await api.buckets.listDataSource();
    if (res.data) {
      console.log(res.data);
      setFiles(res.data);
    }
  };

  const setDbSource = async (file: string) => {
    const res = await api.buckets.changeDataSource(file);
    if (res.status === 200) {
      setSelectedFile(file);
    }
  }

  
  useEffect(() => {
    fetchDbFiles();
  }, []);

  return (
    <div className="w-60 h-full flex flex-col">
      <div className="py-4 flex items-center justify-center gap-2">
        <Database className="h-5 w-5" aria-hidden="true" />
        <h1 className="text-lg">Data</h1>
      </div>
      <ScrollArea className="p-4 h-full w-full">
        <div className="flex flex-col gap-2 mb-20 w-full">
          {files.map((file, index) => (
            <Button
              key={index}
              variant="ghost"
              onClick={() => setDbSource(file)}
              className={`
                ${selectedFile === file
                  ? "bg-primary text-primary-foreground hover:bg-primary/90"
                  : ""} w-full break-words flex flex-wrap h-full
              `}
            >
              <div className="flex flex-col">
                <p className="break-words whitespace-normal">{file.split("/").slice(-3, -1).join(" ")}</p>
                <div className="flex items-center justify-center"><FileIcon className="h-4 w-4 mr-2" aria-hidden="true" /><p className="break-words whitespace-normal">{file.split("/").pop()}</p></div>
              </div>
            </Button>
          ))}
        </div>
      </ScrollArea>
    </div>
);

}
