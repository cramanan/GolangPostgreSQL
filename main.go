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

	log.Fatalln(server.ListenAndServe())
}
