package server

import (
	"fmt"
	"net/http"
	"time"

	"orbiz.one/template-service/src/config"
	"orbiz.one/template-service/src/handler"
	kafka "orbiz.one/template-service/src/kafka/producer"
	"orbiz.one/template-service/src/routes"
	"orbiz.one/template-service/src/service"
)

func Start() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error", err)
		}
	}()

	// Initialize config file
	conf, err := config.GetTemplateConfig()
	if err != nil {
		fmt.Println("error reading config", err)
	}
	// Initialize service
	inventoryService := service.NewInventoryService(
		&config.DBConfig{
			Host:   conf.Postgres.Host,
			Port:   conf.Postgres.Port,
			DBName: conf.Postgres.DBName},
		*kafka.NewKafkaProducer())

	fmt.Println("Service initialized")
	inventoryHandler := handler.InventoryHandler{
		Service: *inventoryService,
	}
	fmt.Println("1")
	server := &http.Server{
		Addr:         ":8081", // http port
		Handler:      routes.SetupRoutes(&inventoryHandler),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	defer func() {
		fmt.Println("connection closed")
		server.Close()
	}()
}
