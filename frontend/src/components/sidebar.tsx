"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { FileIcon } from "@radix-ui/react-icons";
import { useBucketContext } from "@/context/bucket-context";
import { getData, changeDbSource } from "@/api/buckets";
import { Database } from "lucide-react";

export default function Sidebar() {
  const [files, setFiles] = useState<string[]>([]);
  const { selectedFile, setSelectedFile } = useBucketContext();

  const fetchDbFiles = async () => {
    const response = await getData();
    setFiles(response);
  };

  const setDbSource = async (file: string) => {
    console.log(file);
    await changeDbSource(file);
    setSelectedFile(file);
  }

  useEffect(() => {
    fetchDbFiles();
  }, []);

  return (
    <div className="h-full flex flex-col">
      <div className="py-4 flex items-center justify-center gap-2">
        <Database className="h-5 w-5" aria-hidden="true" />
        <h1 className="text-lg">Data</h1>
      </div>
      <ScrollArea className="p-4 h-full">
        <div className="flex flex-col gap-2 mb-20">
          {files.map((file, index) => (
            <Button
              key={index}
              variant="ghost"
              onClick={() => setDbSource(file)}
              className={
                selectedFile === file
                  ? "bg-primary text-primary-foreground hover:bg-primary/90"
                  : ""
              }
            >
              <FileIcon className="h-4 w-4 mr-2" aria-hidden="true" />
              {file.split("/").pop()}
            </Button>
          ))}
        </div>
      </ScrollArea>
    </div>
  );
}
