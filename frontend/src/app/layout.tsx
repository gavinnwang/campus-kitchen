"use client";
// import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

// export const metadata: Metadata = {
//   title: "Create Next App",
//   description: "Generated by create next app",
// };

import { QueryClient, QueryClientProvider } from "react-query";
import { Toaster } from "react-hot-toast";

const queryClient = new QueryClient();

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <html lang="en">
        <QueryClientProvider client={queryClient}>
          <body className={inter.className}>{children}</body>
        </QueryClientProvider>
      </html>
    </>
  );
}

// import { Outlet } from "react-router-dom";

// const queryClient = new QueryClient();
// export default function RootLayout() {
//   return (
//     <QueryClientProvider client={queryClient}>
//       <Outlet />
//     </QueryClientProvider>
//   );
// }
