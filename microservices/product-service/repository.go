package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=product_service sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func CreateProduct(name, description string, price float64) (int32, error) {
	var id int32
	err := db.QueryRow("INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id", name, description, price).Scan(&id)
	return id, err
}

func GetProductByID(id int32) (*Product, error) {
	product := &Product{}
	err := db.QueryRow("SELECT id, name, description, price FROM products WHERE id = $1", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	return product, err
}

func ListProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

type Product struct {
	ID          int32
	Name        string
	Description string
	Price       float64
}
