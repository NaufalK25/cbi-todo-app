package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	Todos = map[int]Todo{
		1: {ID: 1, Title: "Learn Go", Done: false},
		2: {ID: 2, Title: "Build CRUD", Done: false},
		3: {ID: 3, Title: "Connect from Next.js", Done: true},
	}
	NextID = 4
	Mu     sync.Mutex
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
