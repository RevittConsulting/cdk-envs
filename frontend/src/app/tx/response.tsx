"use client";

import { Textarea } from "@/components/ui/textarea";
import { useTxContext } from "@/context/tx-context";

export default function Response() {
  const { response } = useTxContext();
  return (
    <div className="flex flex-col">
      <h1>Response:</h1>
      <Textarea className="resize-none" value={String(response)} disabled />
    </div>
  );
}
