package main

import "github.com/gorilla/websocket"
import "net/http"

type Connection struct {
	ws *websocket.Conn
	// buffered outbound channel
	send chan []byte
}

func (c *Connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()

		if err != nil {
			break
		}
		hub.broadcast <- message
	}
	c.ws.Close()
}

func (c *Connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024,
	WriteBufferSize: 1024}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := &Connection{send: make(chan []byte, 256), ws: ws}
	hub.register <- c
	defer func() { hub.unregister <- c }()
	go c.writer()
	c.reader()
}
