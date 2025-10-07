package main

import (
    "log"
    "net/http"
    "os"

    "github.com/macchiatosoft/go-profiledata-crud-service/database"
    "github.com/macchiatosoft/go-profiledata-crud-service/handler"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	env := &handler.Env{
        DB: db,
        Port: os.Getenv("PORT"),
        Host: os.Getenv("HOST"),
        // We might also have a custom log.Logger, our template instance, and a config struct as fields
    }
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("GET /user/{email}", handler.Handler{env, handler.GetUser})
	log.Println("Server starting on" + env.Port)
	log.Fatal(http.ListenAndServe(":"+env.Port, mux))
}

