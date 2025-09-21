package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/macchiatosoft/go-profiledata-crud-service/database"
)

var db *database.DB

func main() {
	var err error
	db, err = database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{username}", getUserHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	if username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserByUsername(username)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
