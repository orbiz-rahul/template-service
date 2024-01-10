package service

import (
	"fmt"

	"orbiz.one/template-service/src/config"
	"orbiz.one/template-service/src/dal"
	kafka "orbiz.one/template-service/src/kafka/producer"
	"orbiz.one/template-service/src/model"
)

type InventoryService struct {
	config        *config.DBConfig
	dal           dal.InventoryDal
	kafkaProducer kafka.KafkaProducer
}

func NewInventoryService(cfg *config.DBConfig, kafka kafka.KafkaProducer) *InventoryService {

	return &InventoryService{
		config:        cfg,
		kafkaProducer: kafka,
		dal:           dal.GetInventoryDal(cfg),
	}
}

func (s *InventoryService) GetInventoryByID(id int) (*model.Inventory, error) {
	return s.dal.GetByID(id)
}

func (s *InventoryService) CreateInventory(inventory *model.Inventory) error {
	fmt.Println("inside service")
	err := s.dal.Create(inventory)
	if err != nil {
		return err
	}

	// Publish message to Kafka
	s.kafkaProducer.Publish("Inventory Created")

	return nil
}

// Implement UpdateInventory, DeleteInventory methods similarly
