package app

import (
	"fmt"
	"inex/main/configs"
	"log"
)

func init() {
	fmt.Println("run migrate")
	pg, err := configs.NewPG()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pg.URL)
}
