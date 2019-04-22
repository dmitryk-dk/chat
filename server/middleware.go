package server

type Middleware struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (mid *Middleware) Run() {
	for {
		select {
		case client := <-mid.register:
			mid.clients[client] = true
		case client := <-mid.unregister:
			if _, ok := mid.clients[client]; ok {
				delete(mid.clients, client)
				close(client.send)
			}
		case msg := <-mid.broadcast:
			mid.MessageBroadcasting(msg)
		}
	}
}

func (mid *Middleware) MessageBroadcasting(msg []byte) {
	for client := range mid.clients {
		select {
		case client.send <- msg:
		default:
			close(client.send)
			delete(mid.clients, client)
		}
	}
}

func NewMiddleware() *Middleware {
	return &Middleware{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}
