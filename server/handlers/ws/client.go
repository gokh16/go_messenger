<<<<<<< HEAD
package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gokh16/go_messenger/server/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan models.Message
}

func (c *Client) ReadOnConnection() {

	var msg models.Message

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("All clients unregistered")
			panic(err)
		}
		fmt.Println(msg)
		c.hub.broadcast <- msg
	}
}

func (c *Client) WriteOnConnection() {
	defer func() {
		c.conn.Close()
	}()

	for {
		message, ok := <-c.send
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		err := c.conn.WriteJSON(message)
		if err != nil {
			log.Println("Cannot write json")
			panic(err)
		}
	}
}

func ServeWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Cannot upgrade")
		panic(err)
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan models.Message),
	}
	client.hub.register <- client

	go client.WriteOnConnection()
	go client.ReadOnConnection()
}
=======
package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gokh16/go_messenger/server/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan models.Message
}

func (c *Client) ReadOnConnection() {

	var msg models.Message

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("All clients unregistered")
			panic(err)
		}
		fmt.Println(msg)
		c.hub.broadcast <- msg
	}
}

func (c *Client) WriteOnConnection() {
	defer func() {
		c.conn.Close()
	}()

	for {
		message, ok := <-c.send
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		err := c.conn.WriteJSON(message)
		if err != nil {
			log.Println("Cannot write json")
			panic(err)
		}
	}
}

func ServeWebsocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Cannot upgrade")
		panic(err)
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan models.Message),
	}
	client.hub.register <- client

	go client.WriteOnConnection()
	go client.ReadOnConnection()
}
>>>>>>> 3661ec18fda6f6db02155e9be22dd834f0e1cd48
