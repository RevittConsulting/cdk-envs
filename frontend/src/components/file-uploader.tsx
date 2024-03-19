"use client";

import React, { useState } from "react";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import axios from "axios";

export default function FileUploader() {
  const [file, setFile] = useState<File | null>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files) {
      setFile(files[0]);
    }
  };

  const uploadFile = async () => {
    if (file) {
      const formData = new FormData();
      formData.append("file", file);

      const response = await axios.post(
        "https://localhost:8080/api/v1/buckets",
        formData,
        { headers: { "Content-Type": "multipart/form-data" } }
      );
      console.log(response);
    }
  };

  return (
    <div className="flex items-center justify-center flex-row gap-4">
      <Input type="file" onChange={handleFileChange} />
      <Button onClick={uploadFile}>Upload</Button>
    </div>
  );
}
