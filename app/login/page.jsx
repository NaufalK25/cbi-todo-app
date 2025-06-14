"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useFlash } from "../flash-context";

export default function HomePage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { showFlash } = useFlash();

  useEffect(() => {
    const token = window.localStorage.getItem("token");
    if (token) {
      router.push("/");
    }
  }, []);

  const handleLogin = async () => {
    const apiUrl =
      process.env.NODE_ENV === "development"
        ? "http://localhost:3030/api/login"
        : "/api/login";

    const response = await fetch(apiUrl, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
    const data = await response.json();

    if (data.success) {
      window.localStorage.setItem("token", data.data);
      showFlash("Login successfully!", "success");
      setTimeout(() => {
        router.push("/");
      }, 1000);
    } else {
      showFlash(data.error, "error");
    }
  };

  return (
    <main className="flex flex-col gap-2">
      <h1>Login</h1>
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Enter your email"
        className="border border-gray-300 rounded-md py-1 px-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Enter your password"
        className="border border-gray-300 rounded-md py-1 px-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <button
        onClick={handleLogin}
        className="cursor-pointer border py-1 px-2 rounded-md bg-green-700"
      >
        Login
      </button>
    </main>
  );
}
