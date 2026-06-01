import type { Metadata } from "next";
import "leaflet/dist/leaflet.css";

import "@/app/globals.css";


export const metadata: Metadata = {
  title: "Tramplin",
  description:
    "Платформа для стажировок, junior-вакансий, менторских программ и карьерных событий.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ru">
      <body>{children}</body>
    </html>
  );
}
