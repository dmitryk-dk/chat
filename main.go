package main

import (
	"log"

	repository "github.com/dmitryk-dk/chat/repository"
	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewDB("postgres", "postgres://dmitryk:dmitryk@localhost:5432/chat?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer repository.Close(db)

	err = repository.MigrateUp(db)
	if err != nil {
		log.Fatal(err)
	}

}
