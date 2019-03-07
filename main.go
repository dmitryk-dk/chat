package main

import (
	"log"

	repository "github.com/dmitryk-dk/chat/repository"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewDB("postgres", "")
	defer repository.Close(db)
	if err != nil {
		log.Fatal(err)
	}
}
