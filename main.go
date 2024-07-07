package main

import (
	"PostgreSQL/api"
	"log"
)

func main() {
	server, err := api.New(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Starting server")
	log.Fatalln(server.ListenAndServe())
}
