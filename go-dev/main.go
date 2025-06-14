package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	http.HandleFunc("/api/login", withCORS(handleLogin))
	http.HandleFunc("/api/todos", withCORS(handleTodos))

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

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var creds LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid body",
		})
		return
	}

	const validEmail = "admin@example.com"
	const validPassword = "admin1234"

	if creds.Email != validEmail || creds.Password != validPassword {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid credentials",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": creds.Email,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Could not generate token",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Login successfully",
		"data":    tokenString,
	})

}

func handleTodos(w http.ResponseWriter, r *http.Request) {
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
