package server

import (
	"fmt"
	"log"
	"net/http"
)

func ServHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "index.html")
}

func WebsocketConnection(mid *Middleware) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ServeWebSocketConn(mid, w, r)
	}
}

func ServerHttp() {
	http.HandleFunc("/", ServHome)
}

func ServeWebSocketConn(mid *Middleware, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		mid:  mid,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.mid.register <- client

	go client.MessageWrite()
	go client.MessageRead()
}

func ServeWS() {
	mid := NewMiddleware()
	go mid.Run()
	//websocketConn := WebsocketConnection(mid)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWebSocketConn(mid, w, r)
	})
}

func ListenAndServ(addr *string) {
	err := http.ListenAndServe(*addr, nil)
	fmt.Printf("Server start in addr:", &addr)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
