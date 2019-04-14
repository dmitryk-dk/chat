package main

import (
	"flag"

	"github.com/dmitryk-dk/chat/server"

	_ "github.com/lib/pq"
)

var addr = flag.String("addr", ":8080", "http service address") //flag.String("addr", ":8080", "http service address")

func main() {
	// db, err := repository.NewDB("postgres", "postgres://dmitryk:dmitryk@localhost:5432/chat?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer repository.Close(db)

	// err = repository.MigrateUp(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	server.ServerHttp()
	server.ServeWS()
	server.ListenAndServ(addr)
}
