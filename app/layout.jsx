import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { FlashProvider } from "./flash-context";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata = {
  title: "CBI Todo App",
  description: "CBI Todo App (CBI Technical Test for Next.js & Golang)",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <div className="flex items-center justify-center h-screen">
          <FlashProvider>
            <div className="bg-blue-500 text-white p-4 rounded-md shadow">
              {children}
            </div>
          </FlashProvider>
        </div>
      </body>
    </html>
  );
}
