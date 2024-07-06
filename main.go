package main

import (
	"PostgreSQL/api"
	"log"
)

func main() {
	server := api.New(":8080")
	log.Fatalln(server.ListenAndServe())
}
