package routes

import (
	"my-assets-be/controllers"
	"my-assets-be/models"
	"my-assets-be/utils"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nanmu42/etherscan-api"
)

func AccountRoutes(router fiber.Router) {
	etherscanClient := etherscan.NewCustomized(etherscan.Customization{
		Key:     os.Getenv("BSC_SCAN_API_KEY"),
		BaseURL: "https://api.bscscan.com/api?",
	})

	router.Get("/synchronization", func(c *fiber.Ctx) error {
		// Get first transactions of an address.

		txs, err := etherscanClient.NormalTxByAddress(c.Query("address"), nil, nil, 1, 1, false)

		if err != nil {
			panic(err)
		}

		// If address doesn't have any transactions then the synchronization process is completed.
		if len(txs) == 0 {
			return c.JSON(&fiber.Map{
				"status": true,
			})
		}

		prevBlockNumber := txs[0].BlockNumber

		for true {
			blockInfo, err := utils.GetEthBlock(utils.UDecToHex(uint64(prevBlockNumber)))

			if err != nil {
				panic(err)
			}

			if len(blockInfo.Hash) <= 0 {
				break
			}

			address := strings.ToLower(c.Query("address"))
			for _, v := range blockInfo.Transactions {
				if v.To == address {
					num, parseErr := utils.HexToInteger(v.BlockNumber)
					if err := controllers.CreateNormalTransaction(address, models.NormalTransaction{
						Hash:        v.Hash,
						BlockNumber: num,
					}); err != nil || parseErr != nil {
						panic(err)
					}
				}
			}

			prevBlockNumber++
		}

		return c.JSON(&fiber.Map{
			"status": true,
		})
	})

}
