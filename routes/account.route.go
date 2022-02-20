package routes

import (
	"my-assets-be/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(router fiber.Router) {
	// etherscanClient := etherscan.NewCustomized(etherscan.Customization{
	// 	Key:     os.Getenv("BSC_SCAN_API_KEY"),
	// 	BaseURL: "https://api.bscscan.com/api?",
	// })

	router.Get("/synchronization", func(c *fiber.Ctx) error {

		address := strings.ToLower(c.Query("address"))
		utils.GetTransactions(address, 1, 200)
		// Initialize blockNumber with the first found block.
		// prevBlockNumber, err := controllers.GetLatestBlockSync(address)

		// fmt.Println(prevBlockNumber)

		// if err != nil {
		// 	panic(err)
		// }

		// if prevBlockNumber == 0 {
		// 	// Get first transactions of an address.
		// 	txs, err := etherscanClient.NormalTxByAddress(c.Query("address"), nil, nil, 1, 1, false)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	// If address doesn't have any transactions then the synchronization process is completed.
		// 	if len(txs) == 0 {
		// 		return c.JSON(&fiber.Map{
		// 			"status": true,
		// 		})
		// 	}
		// 	prevBlockNumber = int64(txs[0].BlockNumber)
		// }

		// // Looping through the chain until no block has found.
		// for true {
		// 	blockInfo, err := utils.GetEthBlock(utils.UDecToHex(uint64(prevBlockNumber)))
		// 	fmt.Println(prevBlockNumber)

		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	if len(blockInfo.Hash) <= 0 {
		// 		break
		// 	}
		// 	num, parseErr := utils.HexToInteger(blockInfo.Number)
		// 	if parseErr != nil {
		// 		panic(parseErr)
		// 	}
		// 	for _, v := range blockInfo.Transactions {
		// 		if v.To == address || v.From == address {

		// 			if err := controllers.CreateNormalTransaction(address, models.NormalTransaction{
		// 				Hash:        v.Hash,
		// 				BlockNumber: num,
		// 			}); err != nil {
		// 				panic(err)
		// 			}
		// 		}
		// 	}
		// 	if err := controllers.UpdateTransactionSyncStatus(address, num); err != nil {
		// 		panic(err)
		// 	}
		// 	prevBlockNumber += 1
		// }

		return c.JSON(&fiber.Map{
			"status": true,
		})
	})

}
