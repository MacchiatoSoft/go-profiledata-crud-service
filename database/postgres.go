package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/macchiatosoft/go-profiledata-crud-service/types"
)

type DB struct {
	*sql.DB
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

func (db *DB) CreateUser(username string, email string) error {
    var user types.User
    query := "INSERT INTO users (username, email) VALUES ($1, $2)"
    err := db.QueryRow(query, username, email).Scan(&user.ID, &user.Email, &user.Username)
    return err
}

func (db *DB) RemoveUserByEmail(email string) error {
    query := "DELETE FROM users WHERE email = $1"
    _, err := db.Exec(query, email)
    return err
}

func (db *DB) EditUserByEmail(email string, username string, password string) error {
    query := "UPDATE users SET username = $1, password = $2 WHERE email = $3"
    _, err := db.Exec(query, username, password, email)
    return err
}

func (db *DB) GetUserByEmail(email string) (*types.User, error) {
    var user types.User
    query := "SELECT id, email, username FROM users WHERE email = $1"
    err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Username)
    return &user, err
}