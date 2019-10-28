package main

import (
	"gin-restful-demo/routes"
	"log"
)

func main() {
	router := routes.InitRouter()
	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)

	}
}
