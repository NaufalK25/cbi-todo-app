// app/flash-context.js
"use client";

import { createContext, useContext, useState } from "react";

const FlashContext = createContext();

export const FlashProvider = ({ children }) => {
  const [message, setMessage] = useState(null);

  const showFlash = (msg, type = "error", duration = 3000) => {
    setMessage({ text: msg, type });

    setTimeout(() => {
      setMessage(null);
    }, duration);
  };

  return (
    <FlashContext.Provider value={{ showFlash }}>
      {message && (
        <div
          className={`fixed top-4 left-1/2 -translate-x-1/2 px-4 py-2 rounded shadow z-50 text-white ${
            message.type === "error" ? "bg-red-700" : "bg-green-700"
          }`}
        >
          {message.text}
        </div>
      )}
      {children}
    </FlashContext.Provider>
  );
};

export const useFlash = () => useContext(FlashContext);
