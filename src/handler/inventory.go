package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"orbiz.one/template-service/src/service"

	"orbiz.one/template-service/src/model"
)

type InventoryHandler struct {
	Service service.InventoryService
	//  Validator middleware.ValidateJWT // Your validation logic implementation
	// Other dependencies like JWT authentication, Kafka producer, etc.
}

func (h *InventoryHandler) CreateInventory(w http.ResponseWriter, r *http.Request) {
	// Parse request body into Inventory struct
	var newInventory model.Inventory
	if err := json.NewDecoder(r.Body).Decode(&newInventory); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the incoming data
	//if err := h.Validator.Validate(newInventory); err != nil {
	//    http.Error(w, "Validation failed", http.StatusBadRequest)
	//    return
	//}

	// Perform JWT authentication
	// if !h.authenticate(r) {
	//     http.Error(w, "Unauthorized", http.StatusUnauthorized)
	//     return
	// }
	fmt.Println("inside handler")
	// Create the inventory item
	if err := h.Service.CreateInventory(&newInventory); err != nil {
		http.Error(w, "Failed to create inventory", http.StatusInternalServerError)
		return
	}

	// Produce Kafka message for the created inventory
	// h.KafkaProducer.Produce("Inventory created: " + newInventory.Name)

	// Set HTTP status code and respond with the created inventory
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newInventory)
}

// Implement other handler functions for GET, PUT, DELETE operations similarly

// Your JWT authentication logic goes here

// Your Kafka producer logic goes here
