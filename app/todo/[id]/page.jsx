"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { useRouter, useParams } from "next/navigation";

export default function TodoDetailPage() {
  const router = useRouter();
  const { id } = useParams();
  const [todo, setTodo] = useState(null);
  const [loading, setLoading] = useState(true);

  const apiUrl =
    process.env.NODE_ENV === "development"
      ? `http://localhost:3030/api/todo?id=${id}`
      : `/api/todo?id=${id}`;

  useEffect(() => {
    fetch(apiUrl)
      .then((res) => res.json())
      .then(setTodo)
      .catch(console.error)
      .finally(() => setLoading(false));
  }, [apiUrl]);

  const updateTodo = async () => {
    if (!todo) return;
    await fetch(apiUrl, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(todo),
    });
    router.push("/");
  };

  const deleteTodo = async () => {
    await fetch(apiUrl, { method: "DELETE" });
    router.push("/");
  };

  if (loading) return <p>Loading...</p>;
  if (!todo) return <p>Todo not found.</p>;

  return (
    <main className="flex flex-col gap-2">
      <h1>Todo Detail</h1>
      <input
        value={todo.title}
        onChange={(e) => setTodo({ ...todo, title: e.target.value })}
        className="border border-gray-300 rounded-md py-1 px-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <label className="inline-flex items-center space-x-2">
        <input
          type="checkbox"
          checked={todo.done}
          onChange={(e) => setTodo({ ...todo, done: e.target.checked })}
        />
        Done
      </label>
      <div className="flex gap-2">
        <Link href="/" className="border py-1 px-2 rounded-md">
          Back
        </Link>
        <button
          onClick={updateTodo}
          className="cursor-pointer border py-1 px-2 rounded-md bg-yellow-700"
        >
          Update
        </button>
        <button
          onClick={deleteTodo}
          className="cursor-pointer border py-1 px-2 rounded-md bg-red-700"
        >
          Delete
        </button>
      </div>
    </main>
  );
}
