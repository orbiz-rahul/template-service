package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"orbiz.one/template-service/src/model"
)

type InventoryRepo struct {
	db *sql.DB
}

func NewInventoryRepository() *InventoryRepo {
	//var db *sql.DB
	db, err := sql.Open("postgres", "http://localhost/postgrace?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DB connected")
	return &InventoryRepo{db: db}
}

func (r *InventoryRepo) GetByID(id int) (*model.Inventory, error) {
	// Fetch inventory details from the database
	// Implement logic to fetch data from the database using SQL queries
	// return Inventory object or error
	return nil, nil
}

func (r *InventoryRepo) Create(inventory *model.Inventory) error {
	// Insert inventory data into the database
	// Implement logic to insert data into the database using SQL queries
	// return error if any
	return nil
}

// Implement Update and Delete methods similarly
