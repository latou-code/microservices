package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=inventory_service sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetAvailableStock(productID int32) (int32, error) {
	var quantity int32
	err := db.QueryRow("SELECT quantity FROM inventory WHERE product_id = $1", productID).Scan(&quantity)
	if err != nil {
		return 0, err
	}
	return quantity, nil
}

func DecreaseStock(productID int32, quantity int32) (bool, error) {
	res, err := db.Exec("UPDATE inventory SET quantity = quantity - $1 WHERE product_id = $2 AND quantity >= $1", quantity, productID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
