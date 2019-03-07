package main

import (
	"log"

	"github.com/dmitryk-dk/chat/storage"

	_ "github.com/lib/pq"
)

func main() {
	_, err := storage.NewDB("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
}
