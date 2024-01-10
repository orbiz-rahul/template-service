package model

type Inventory struct {
	ID       int
	Name     string
	Quantity int
	// Other inventory-related fields
}

type InventoryService interface {
	GetInventoryByID(id int) (*Inventory, error)
	CreateInventory(inventory *Inventory) error
	UpdateInventory(inventory *Inventory) error
	DeleteInventory(id int) error
}

type InventoryDal interface {
	GetByID(id int) (*Inventory, error)
	Create(inventory *Inventory) error
	Update(inventory *Inventory) error
	Delete(id int) error
}
