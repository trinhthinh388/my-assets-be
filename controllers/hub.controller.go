package controllers

import (
	"log"
	"my-assets-be/types"
	"strings"
	"sync"

	"github.com/gofiber/websocket/v2"
)

var WSSync sync.Once
var clients map[string]types.Client
var Broadcast = make(chan types.WSMessage)
var Register = make(chan *websocket.Conn)
var Unregister = make(chan *websocket.Conn)

func GetWSClients() map[string]types.Client {
	WSSync.Do(func() {
		clients = make(map[string]types.Client)
	})

	return clients
}

func RunHub() {
	clients := GetWSClients()
	for {
		select {
		case connection := <-Register:
			clients[strings.ToLower(connection.Params("id"))] = types.Client{
				Conn: connection,
			}
			log.Println("connection registered")

		case message := <-Broadcast:
			log.Println("message received:", message.Content)
			log.Println("message received from:", message.From)

			// Send the message to an specific clients
			client := clients[message.From]
			if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message.Content)); err != nil {
				log.Println("write error:", err)

				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				client.Conn.Close()
				delete(clients, message.From)
			}

		case connection := <-Unregister:
			// Remove the client from the hub
			delete(clients, connection.Params("id"))

			log.Println("connection unregistered")
		}
	}
}

func SendTextMessage(id string, message string) {
	clients := GetWSClients()
	client := clients[id]
	if client.Conn != nil {
		if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("write error:", err)

			client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			client.Conn.Close()
			delete(clients, id)
		}
	}
}
