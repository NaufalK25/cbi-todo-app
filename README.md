# CBI Todo App

## Prerequisites

1. [Node.js](https://nodejs.org/en/download)
   ```bash
   php --version
   npm --version
   ```
2. [Go](https://go.dev/doc/install)
   ```bash
   go version
   ```
3. [Git](https://git-scm.com/downloads)
   ```bash
   git --version
   ```

## How To Run In Local

1. Clone the repository
   ```bash
   git clone https://github.com/NaufalK25/cbi-todo-app.git
   ```
2. Install dependencies
   ```bash
   npm i
   ```
3. Copy .env.example
   ```bash
   cp .env.example .env
   ```
4. Download Go modules
   ```bash
   go mod download
   ```
5. Run the server
   ```bash
    go run go-dev/main.go
   ```
6. Run the client
   ```
   npm run dev
   ```
