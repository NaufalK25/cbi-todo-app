"use client";

import { useEffect, useState } from "react";
import Link from "next/link";

export default function HomePage() {
  const [todos, setTodos] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const url =
      process.env.NODE_ENV === "development"
        ? "http://localhost:3030/api/todos"
        : "/api/todos";

    fetch(url)
      .then((res) => res.json())
      .then(setTodos)
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, []);

  return (
    <main className="flex flex-col gap-2">
      <h1>All Todos</h1>
      <Link
        href="/todo/new"
        className="border py-1 px-2 rounded-md bg-emerald-700 w-fit"
      >
        + Add Todo
      </Link>
      {loading ? (
        <p>Loading...</p>
      ) : todos.length === 0 ? (
        <p>No todos yet</p>
      ) : (
        <ul className="flex flex-col gap-2">
          {todos.map((todo) => (
            <li key={todo.id}>
              <Link href={`/todo/${todo.id}`}>
                {todo.done ? "✅" : "❎"} {todo.title}
              </Link>
            </li>
          ))}
        </ul>
      )}
    </main>
  );
}
