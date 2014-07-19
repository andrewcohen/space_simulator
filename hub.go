package main

type Hub struct {
	connections map[*Connection]bool
	// inbound messages
	broadcast chan []byte

	register   chan *Connection
	unregister chan *Connection
}

var hub = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func (hub *Hub) run() {
	for {
		select {
		case c := <-hub.register:
			hub.connections[c] = true
		case c := <-hub.unregister:
			if _, ok := hub.connections[c]; ok {
				delete(hub.connections, c)
				close(c.send)
			}
		case m := <-hub.broadcast:
			for c := range hub.connections {
				select {
				case c.send <- m:
				default:
					delete(hub.connections, c)
					close(c.send)
				}
			}
		}
	}
}
