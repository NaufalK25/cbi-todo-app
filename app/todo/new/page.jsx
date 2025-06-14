"use client";

import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function CreateTodo() {
  const [title, setTitle] = useState("");
  const router = useRouter();

  const handleCreate = async () => {
    const url =
      process.env.NODE_ENV === "development"
        ? "http://localhost:3030/api/todos"
        : "/api/todos";

    if (!title) {
      return alert("Title can't be empty!");
    }

    await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title, done: false }),
    });

    router.push("/");
  };

  return (
    <main className="flex flex-col gap-2">
      <h1>Create Todo</h1>
      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Enter todo title"
        className="border border-gray-300 rounded-md py-1 px-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <div className="flex gap-2">
        <Link href="/" className="border py-1 px-2 rounded-md">
          Back
        </Link>
        <button
          onClick={handleCreate}
          className="cursor-pointer border py-1 px-2 rounded-md bg-green-700"
        >
          Create
        </button>
      </div>
    </main>
  );
}
