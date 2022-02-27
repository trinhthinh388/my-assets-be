package routes

import (
	"my-assets-be/controllers"
	"my-assets-be/database"
	"my-assets-be/types"
	"my-assets-be/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(router fiber.Router) {
	router.Get("/stats", func(c *fiber.Ctx) error {
		address := strings.ToLower(c.Query("address"))
		chain := strings.ToLower(c.Query("chain"))
		go controllers.AnalyzeStats(address, chain)
		return c.JSON(&fiber.Map{
			"status": true,
		})
	})

	router.Get("/synchronization", func(c *fiber.Ctx) error {
		address := strings.ToLower(c.Query("address"))
		chain := strings.ToLower(c.Query("chain"))

		if utils.IsValidAddress(address) {
			return c.JSON(&fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "address is in wrong format",
			})
		}

		if chain == "" {
			return c.JSON(&fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "wrong chain",
			})
		}

		go controllers.SyncTx(address, chain)

		return c.JSON(&fiber.Map{
			"code":   200,
			"status": true,
		})
	})

	router.Get("/transactions", func(c *fiber.Ctx) error {
		address := strings.ToLower(c.Query("address"))
		fromBlock, _ := strconv.ParseInt(c.Query("fromBlock"), 10, 64)
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		desc, _ := strconv.ParseBool(c.Query("desc"))
		startValue := types.TxDetail{
			BlockNumber: int(fromBlock),
		}
		txs, err := database.GetTx(address, startValue, int(limit), desc)

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"result": txs,
			"total":  len(txs),
		})
	})

}
