package service

import (
	"fmt"

	"orbiz.one/template-service/src/dal/db"
	kafka "orbiz.one/template-service/src/kafka/producer"
	"orbiz.one/template-service/src/model"
)

type InventoryService struct {
	repository    db.InventoryRepo
	kafkaProducer kafka.KafkaProducer
}

func NewInventoryService(repo db.InventoryRepo, kafka kafka.KafkaProducer) *InventoryService {
	return &InventoryService{
		repository:    repo,
		kafkaProducer: kafka,
	}
}

func (s *InventoryService) GetInventoryByID(id int) (*model.Inventory, error) {
	return s.repository.GetByID(id)
}

func (s *InventoryService) CreateInventory(inventory *model.Inventory) error {
	fmt.Println("inside service")
	err := s.repository.Create(inventory)
	if err != nil {
		return err
	}

	// Publish message to Kafka
	s.kafkaProducer.Publish("Inventory Created")

	return nil
}

// Implement UpdateInventory, DeleteInventory methods similarly
