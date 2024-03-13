import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import Nav from "@/components/nav";
import { BucketProvider } from "@/context/bucket-context";
import { ChainProvider } from "@/context/chain-context";
import { TxProvider } from "@/context/tx-context";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "cdk-envs",
  description: "cdk-envs",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={inter.className}>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem={true}
          themes={["light", "dark"]}
          disableTransitionOnChange
        >
          <BucketProvider>
          <ChainProvider>
          <TxProvider>
            <Nav />
            {children}
          </TxProvider>
          </ChainProvider>
          </BucketProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
