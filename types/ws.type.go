package types

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Clients struct {
	Clients map[string]Client
	mu      sync.Mutex
}

func (c *Clients) Add(key string, value *Client) {
	c.mu.Lock()
	c.Clients[key] = *value
	c.mu.Unlock()
}

func (c *Clients) GetClient(key string) Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Clients[key]
}

func (c *Clients) DeleteClient(key string) {
	c.mu.Lock()
	delete(c.Clients, key)
	c.mu.Unlock()
}

type Client struct {
	// The websocket connection.
	Conn *websocket.Conn
}

type WSMessage struct {
	From    string
	Content string
}
