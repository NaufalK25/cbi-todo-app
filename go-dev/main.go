package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	Todos  map[int]Todo = map[int]Todo{}
	Mu     sync.Mutex
	NextID int = 1
)

func main() {
	seedData()

	http.HandleFunc("/api/todos", withCORS(handleTodos))
	http.HandleFunc("/api/todo", withCORS(handleTodo))

	fmt.Println("Server running at http://localhost:3030")
	http.ListenAndServe(":3030", nil)
}

func seedData() {
	Todos[NextID] = Todo{ID: NextID, Title: "Learn Go", Done: false}
	NextID++
	Todos[NextID] = Todo{ID: NextID, Title: "Build CRUD", Done: false}
	NextID++
	Todos[NextID] = Todo{ID: NextID, Title: "Connect from Next.js", Done: true}
	NextID++
}

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	}
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	Mu.Lock()
	defer Mu.Unlock()

	switch r.Method {
	case "GET":
		todos := []Todo{}
		for _, t := range Todos {
			todos = append(todos, t)
		}
		json.NewEncoder(w).Encode(todos)

	case "POST":
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}
		todo.ID = NextID
		NextID++
		Todos[todo.ID] = todo
		json.NewEncoder(w).Encode(todo)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	Mu.Lock()
	defer Mu.Unlock()

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id == 0 {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	todo, exists := Todos[id]
	if !exists {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(todo)

	case "PUT":
		var updated Todo
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}
		updated.ID = id
		Todos[id] = updated
		json.NewEncoder(w).Encode(updated)

	case "DELETE":
		delete(Todos, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
