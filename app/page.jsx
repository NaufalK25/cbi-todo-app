"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useFlash } from "./flash-context";

export default function HomePage() {
  const router = useRouter();
  const [todos, setTodos] = useState([]);
  const [loading, setLoading] = useState(true);
  const { showFlash } = useFlash();

  useEffect(() => {
    const token = window.localStorage.getItem("token");
    if (!token) {
      router.push("/login");
    }
  }, []);

  useEffect(() => {
    const fetchTodos = async () => {
      const apiUrl =
        process.env.NODE_ENV === "development"
          ? "http://localhost:3030/api/todos"
          : "/api/todos";

      const response = await fetch(apiUrl);
      const data = await response.json();

      if (data.success) {
        showFlash(data.message, "success");
        setTodos(data.data);
      } else {
        showFlash(data.error, "error");
      }

      setLoading(false);
    };

    fetchTodos();
  }, []);

  const handleLogout = () => {
    window.localStorage.removeItem("token");
    showFlash("Logout successfully!", "success");
      setTimeout(() => {
        router.push("login");
      }, 1000);
  };

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
      <button
        onClick={handleLogout}
        className="cursor-pointer border py-1 px-2 rounded-md bg-red-700 w-fit"
      >
        Logout
      </button>
    </main>
  );
}
