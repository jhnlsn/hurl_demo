package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	mu    sync.Mutex
	items = []Item{
		{ID: 1, Name: "apple"},
		{ID: 2, Name: "banana"},
	}
	nextID = 3
)

func main() {
	http.HandleFunc("GET /health", handleHealth)
	http.HandleFunc("GET /items", handleListItems)
	http.HandleFunc("GET /items/{id}", handleGetItem)
	http.HandleFunc("POST /items", handleCreateItem)
	http.HandleFunc("DELETE /items/{id}", handleDeleteItem)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func handleListItems(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func handleGetItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Name == "" {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	mu.Lock()
	item := Item{ID: nextID, Name: body.Name}
	items = append(items, item)
	nextID++
	mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
