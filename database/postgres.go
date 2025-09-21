package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func Connect() (*DB, error) {
	dsn := fmt.Sprintf("host=db port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec("SET search_path TO auth, data, public"); err != nil {
		return nil, err
	}

	return &DB{db}, err
}

func (db *DB) GetUserByUsername(username string) (*User, error) {
	var user User
	query := "SELECT id, email, username FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Email, &user.Username)
	return &user, err
}
