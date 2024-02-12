"use client";

import Link from "next/link";
import Image from "next/image";
import { ModeToggle } from "@/components/mode-toggle";

const Nav = () => {
  return (
    <nav className="flex items-center justify-between p-2 mx-4 md:mx-32">
      <div className="flex items-center">
        <Link passHref href="/">
          <Image
            src="/gateway-logo.svg"
            alt="logo"
            width={50}
            height={50}
            className="object-contain rounded-lg mx-auto height-auto"
            priority
          />
        </Link>
      </div>

      <div className="flex justify-center items-center gap-4">
        <ModeToggle />
      </div>
    </nav>
  );
};

export default Nav;