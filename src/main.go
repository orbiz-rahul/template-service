package main

import (
	"fmt"

	"orbiz.one/template-service/src/server"
)

func main() {
	fmt.Println("IMS started...")

	server.Start()
}
