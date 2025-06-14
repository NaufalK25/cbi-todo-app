package api

import "sync"

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
