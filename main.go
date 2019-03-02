package main

import (
	"log"

	"github.com/dmitryk-dk/chat/storage"
)

func main() {
	dbStore, err := storage.New("mysql", "someceredentials")
	if err != nil {
		log.Fatal(err)
	}
}
