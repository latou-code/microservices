package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=cart_service sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func AddProductToCart(userID, productID, quantity int32) error {
	_, err := db.Exec("INSERT INTO cart (user_id, product_id, quantity) VALUES ($1, $2, $3)", userID, productID, quantity)
	return err
}

func GetCartItems(userID int32) ([]CartItem, error) {
	rows, err := db.Query("SELECT product_id, quantity FROM cart WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CartItem
	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ProductID, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

type CartItem struct {
	ProductID int32
	Quantity  int32
}
