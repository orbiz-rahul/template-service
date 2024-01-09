package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"orbiz.one/template-service/src/handler"
)

func SetupRoutes(inventoryHandler *handler.InventoryHandler) *http.ServeMux {

	fmt.Println("handler invoked")
	routers := mux.NewRouter()
	// Create routes and associate handlers
	routers.HandleFunc("/inventory", inventoryHandler.CreateInventory).Methods(http.MethodPost)
	httpRouter := http.NewServeMux()
	httpRouter.Handle("/", routers)
	return httpRouter
	// Add other routes and handlers for different endpoints

	// For example:
	// http.HandleFunc("/inventory/update", inventoryHandler.UpdateInventory)
	// http.HandleFunc("/inventory/delete", inventoryHandler.DeleteInventory)
	// ...
}
