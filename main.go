package main

import (
	"log"

	"github.com/dmitryk-dk/chat/storage"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := storage.NewDB("mysql", "dmitryk:dmitryk@tcp(localhost:3306)/chat")
	if err != nil {
		log.Fatal(err)
	}
}
