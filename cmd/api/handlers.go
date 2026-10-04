package main

import (
	"encoding/json"
	"net/http"
)
// list all registered routes
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routes)
}

// get all users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	if err := db.NewSelect().Model(&users).Order("id").Scan(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid body JSON", http.StatusBadRequest)
		return
	}
	u.ID = 0
	if _, err := db.NewInsert().Model(&u).Exec(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)

}

// curl -X POST localhost:8080/users -d '{"name":"Dennis"}'
