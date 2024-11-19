package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=order_service sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func CreateOrder(userID int32) (int32, error) {
	var orderID int32
	err := db.QueryRow("INSERT INTO orders (user_id) VALUES ($1) RETURNING id", userID).Scan(&orderID)
	return orderID, err
}

type Order struct {
	ID     int32
	UserID int32
}
