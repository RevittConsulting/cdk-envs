import TxForm from "@/components/tx/tx-form";
import Response from "./response";

export default function TxPage() {
  return (
    <div className="flex flex-col items-center justify-center gap-4 container py-10">
      <h1 className="font-bold text-3xl">Tx Sender</h1>

      <div className="flex justify-center items-start w-full gap-4">
        <TxForm />

        <div className="w-[30vw]">
          <Response />
        </div>
      </div>
    </div>
  );
}
