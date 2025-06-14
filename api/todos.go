package handler

import (
	"encoding/json"
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

	idParam := r.URL.Query().Get("id")

	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil || id == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid id",
			})
			return
		}

		todo, exists := Todos[id]
		if !exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Todo not found",
			})
			return
		}

		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"message": "Todo fetched successfully",
				"data":    todo,
			})
			return

		case "PUT":
			var updated Todo
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"error":   "Invalid body",
				})
				return
			}
			updated.ID = id
			Todos[id] = updated
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"message": "Todo updated successfully",
				"data":    updated,
			})
			return

		case "DELETE":
			delete(Todos, id)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"message": "Todo deleted successfully",
				"data":    nil,
			})
			return

		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Method not allowed",
			})
			return
		}
	}

	switch r.Method {
	case "GET":
		todos := []Todo{}
		for _, t := range Todos {
			todos = append(todos, t)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "Todos fetched successfully",
			"data":    todos,
		})

	case "POST":
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid body",
			})
			return
		}
		todo.ID = NextID
		NextID++
		Todos[todo.ID] = todo
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "Todo created successfully",
			"data":    todo,
		})

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}
}
