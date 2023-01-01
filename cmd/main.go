package main

import (
	"fmt"
	routes "inex/main/internal/transport/routes"
	"log"
	"net/http"
	// routes "inex/main/routes"
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