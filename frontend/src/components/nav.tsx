"use client";

import Link from "next/link";
import { ModeToggle } from "@/components/mode-toggle";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

const Nav = () => {
  const router = useRouter();
  const visitDataview = () => {
    router.push("/dataview");
  };
  const visitDatastream = () => {
    router.push("/datastream");
  };
  const visitTx = () => {
    router.push("/tx");
  };
  return (
    <nav className="flex items-center justify-between p-2 mx-4 md:mx-32">
      <div className="flex items-center gap-4">
        <Link passHref href="/">
          <span className="font-extrabold text-3xl text-primary">CDU</span>
        </Link>
        <Button variant="outlineprimary" onClick={() => {visitDataview()}}>
          Dataview
        </Button>
        <Button variant="outlineprimary" onClick={() => {visitDatastream()}}>
          Datastream
        </Button>
        <Button variant="outlineprimary" onClick={() => {visitTx()}}>
          Tx
        </Button>
      </div>

      <div className="flex justify-center items-center gap-4">
        <ModeToggle />
      </div>
    </nav>
  );
};

export default Nav;