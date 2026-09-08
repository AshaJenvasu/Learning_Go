package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	dbname = "test"
	user = "postgres"
	password = "mypassword"
)

var db *sql.DB

type Product struct {
	ID int
	Name string
	Price int
}

func main() {
	// Connection string
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  // Open a connection
  sdb, err := sql.Open("postgres", psqlInfo)

  if err != nil {
    log.Fatal(err)
  }

	db = sdb

  defer db.Close()

  // Check the connection
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Successfully connected!")

product, err := getProduct(2)
if err != nil {
    log.Fatal(err)
  }
	fmt.Println("Get Successful !", product)
}

func createProduct(product *Product) error {
_, err := db.Exec(
	"INSERT INTO public.products(name, price)VALUES ($1, $2);",
	product.Name,
	product.Price,
)

	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow(
		"SELECT id,name,price FROM products WHERE id=$1;",
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err!= nil {
		return Product{}, err
	}

	return p, nil
}