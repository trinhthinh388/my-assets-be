package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"

	"my-assets-be/config"
	"my-assets-be/controllers"
	"my-assets-be/routes"
	"my-assets-be/types"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error while loading .env file.")
	}

	config.Migrate()

	app := fiber.New()

	tracker, err := controllers.NewTracker("https://data-seed-prebsc-1-s1.binance.org:8545/", "BNB Testnet")

	if err != nil {
		panic(err)
	}

	tracker.Track()

	go controllers.RunHub()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		defer func() {
			controllers.Unregister <- c
			c.Close()
		}()

		controllers.Register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				payload := types.WSMessage{
					From:    c.Params("id"),
					Content: string(message),
				}
				// Broadcast the received message
				controllers.Broadcast <- payload
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	v1 := app.Group("/v1")

	routes.AccountRoutes(v1)

	app.Listen(":3000")
}
