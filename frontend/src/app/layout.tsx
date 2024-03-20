import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import Nav from "@/components/nav";
import { BucketProvider } from "@/context/bucket-context";
import { ChainProvider } from "@/context/chain-context";
import { TxProvider } from "@/context/tx-context";
import { TooltipProvider } from "@/components/ui/tooltip";

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
                <TooltipProvider>
                  <div className="flex flex-col h-full">
                    <Nav />
                    <main className="fex-1 overflow-auto h-full">
                      {children}
                    </main>
                  </div>
                </TooltipProvider>
              </TxProvider>
            </ChainProvider>
          </BucketProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
