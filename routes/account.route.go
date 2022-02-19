package routes

import (
	"fmt"
	"my-assets-be/controllers"
	"my-assets-be/geth"
	"my-assets-be/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nanmu42/etherscan-api"
)

const (
	minOffset int64 = 1000
)

func AccountRoutes(router fiber.Router) {
	etherscanClient := etherscan.NewCustomized(etherscan.Customization{
		Key:     os.Getenv("BSC_SCAN_API_KEY"),
		BaseURL: "https://api.bscscan.com/api?",
	})

	router.Get("/synchronization", func(c *fiber.Ctx) error {
		count, err := controllers.GetNormalTransactionCount(c.Query("address"))
		if err != nil {
			panic(err)
		}

		firstTx, err := etherscanClient.NormalTxByAddress(c.Query("address"), nil, nil, 1, 1, false)

		if err != nil {
			panic(err)
		}

		fmt.Println(firstTx[0].BlockNumber)

		return c.JSON(&fiber.Map{
			"count": count,
		})
	})

	router.Get("/transactions", func(c *fiber.Ctx) error {
		txs, err := etherscanClient.NormalTxByAddress(c.Query("address"), nil, nil, 2, 10, false)

		if err != nil {
			panic(err)
		}

		tx, err := controllers.GetLatestNormalTransaction(c.Query("address"))

		if err != nil {
			panic(err)
		}

		fmt.Println(tx.Hash)

		// list, err := controllers.GetAllNormalTransaction(c.Query("address"))

		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Println(list)

		// normalTxs := make([]models.NormalTransaction, len(txs))

		// for i, v := range txs {
		// 	normalTxs[i] = models.NormalTransaction{
		// 		BlockNumber:       v.BlockNumber,
		// 		TimeStamp:         models.Time(v.TimeStamp),
		// 		Hash:              v.Hash,
		// 		Nonce:             v.Nonce,
		// 		BlockHash:         v.BlockHash,
		// 		TransactionIndex:  v.TransactionIndex,
		// 		From:              v.From,
		// 		To:                v.To,
		// 		Value:             (*models.BigInt)(v.Value),
		// 		Gas:               v.Gas,
		// 		GasPrice:          (*models.BigInt)(v.GasPrice),
		// 		IsError:           v.IsError,
		// 		TxReceiptStatus:   v.TxReceiptStatus,
		// 		Input:             v.Input,
		// 		ContractAddress:   v.ContractAddress,
		// 		CumulativeGasUsed: v.CumulativeGasUsed,
		// 		GasUsed:           v.GasUsed,
		// 		Confirmations:     v.Confirmations,
		// 	}
		// }

		// if err := controllers.CreateNormalTransactions(c.Query("address"), normalTxs); err != nil {
		// 	panic(err)
		// }

		logs := make(map[string][]geth.TxReceiptLog)

		for i := 0; i < len(txs); i++ {
			receipt, err := geth.GetTransactionReceipt(txs[i].Hash)
			if err != nil {
				continue
			}
			logs[txs[i].Hash] = utils.FilterTransferLogs(receipt.Logs)
		}

		return c.JSON(&fiber.Map{
			"transactions": txs,
			"logs":         logs,
		})
	})

}
