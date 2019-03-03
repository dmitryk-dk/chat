package main

import (
	"log"
	"time"

	"github.com/dmitryk-dk/chat/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := models.NewDB("mysql", "dmitryk:dmitryk@tcp(localhost:3306)/chat")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	dateTime := time.Now().Local().Format("2006-01-02 15:04:05")
	user := models.User{10, "petya", dateTime}
	err = user.Create(db.DB)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)
	msg := models.Message{27, "text from user 10", dateTime}
	err = msg.Create(db.DB)
	if err != nil {
		log.Fatal(err)
	}
}
