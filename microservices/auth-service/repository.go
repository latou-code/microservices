package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=auth_service sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func CreateUser(username, email, password string) error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
	return err
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := db.QueryRow("SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
	return user, err
}

type User struct {
	Email    string
	Password string
}
