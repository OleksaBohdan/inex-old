package main

import (
	"fmt"
	"inex/main/internal/controller/http/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello inex")

	routes.HttpRoutes()

	fmt.Println("Server started")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
