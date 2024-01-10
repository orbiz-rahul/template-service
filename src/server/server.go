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

type listenAndServerFunc = func(addr string, handler http.Handler)
type ServerImpl struct {
	ListenAndServer listenAndServerFunc
}

func Start() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error", err)
		}
	}()
	fmt.Println("Starting service")
	inventoryService := service.NewInventoryService(&config.DBConfig{Host: "localhost", Port: "5432", DBName: "postgres"}, *kafka.NewKafkaProducer())
	fmt.Println("Service initialised")
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
