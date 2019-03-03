package main

import (
	"log"
	"time"

	"github.com/dmitryk-dk/chat/models"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var id uuid.UUID
	var user models.User
	db, err := models.NewDB("mysql", "dmitryk:dmitryk@tcp(localhost:3306)/chat")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	user, err = db.GetUser("0974b2c7-9c22-4c7c-b6cf-65a19fce2028")
	if err != nil {
		log.Fatal(err)
	}

	dateTime := time.Now().Local().Format("2006-01-02 15:04:05")
	id, err = uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	user = models.User{id, "petya", dateTime}
	err = user.Create(db.DB)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)
	msg := models.Message{user.ID, "text from user 10", dateTime}
	err = msg.Create(db.DB)
	if err != nil {
		log.Fatal(err)
	}
}
