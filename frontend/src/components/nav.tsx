"use client";

import Link from "next/link";
import Image from "next/image";
import { ModeToggle } from "@/components/mode-toggle";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

const Nav = () => {
  const router = useRouter();
  const visitDatacryp = () => {
    router.push("/datacryp");
  };
  return (
    <nav className="flex items-center justify-between p-2 mx-4 md:mx-32">
      <div className="flex items-center gap-4">
        <Link passHref href="/">
          <Image
            src="/gateway-logo.svg"
            alt="logo"
            width={40}
            height={40}
            priority
          />
        </Link>
        <Button variant="outlineprimary" onClick={() => {visitDatacryp()}}>
          Datacryp
        </Button>
      </div>

      <div className="flex justify-center items-center gap-4">
        <ModeToggle />
      </div>
    </nav>
  );
};

export default Nav;