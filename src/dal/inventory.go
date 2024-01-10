package dal

import (
	"fmt"

	_ "github.com/lib/pq"
	"orbiz.one/template-service/src/config"
	"orbiz.one/template-service/src/model"
)

type InventoryDal struct {
	config *config.DBConfig
}

func (r *InventoryDal) GetByID(id int) (*model.Inventory, error) {
	// Fetch inventory details from the database
	// Implement logic to fetch data from the database using SQL queries
	// return Inventory object or error
	return nil, nil
}

func (r *InventoryDal) Create(inventory *model.Inventory) error {

	fmt.Println("inside dal")
	mgr := GetDBManager(r.config)
	db, err := mgr.GetDBConnector()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO items (id, name, quantity) VALUES ($1, $2, $3)", inventory.ID, inventory.Name, inventory.Quantity)
	if err != nil {
		fmt.Println("error while creating record: ", err)
		return err
	}
	fmt.Println("Item inserted successfully!")

	return nil
}

func GetInventoryDal(cfg *config.DBConfig) InventoryDal {
	return InventoryDal{
		config: cfg,
	}
}
