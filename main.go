package main

import (
	"log"
	"time"

	"github.com/dmitryk-dk/chat/storage"
	"github.com/dmitryk-dk/chat/user"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbStore, err := storage.New("mysql", "dmitryk:dmitryk@tcp(localhost:3306)/chat")
	defer dbStore.DB.Close()
	if err != nil {
		log.Fatal(err)
	}

	dateTime := time.Now().Local().Format("2006-01-02 15:04:05")
	user := user.User{1, "vasya", dateTime, dbStore}
	err = user.Create()
	if err != nil {
		log.Fatal(err)
	}
}
