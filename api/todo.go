package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"cbi-todo-app/api/lib/store"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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