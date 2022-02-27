package controllers

import (
	"fmt"
	"math/big"
	"my-assets-be/database"
	"my-assets-be/types"
	"my-assets-be/utils"
)

func AnalyzeStats(address string, chain string) {
	var txs []types.TxDetail
	var startValue types.TxDetail

	stats, err := database.GetStat(address, chain)
	if err != nil {
		panic(err)
	}

	if stats == nil {
		startValue = types.TxDetail{
			BlockNumber: 0,
		}
		stats = &types.AddressStats{
			Address:    address,
			BuyAmount:  types.BigInt{},
			SellAmount: types.BigInt{},
		}
	} else {
		startValue = types.TxDetail{
			BlockNumber: int(stats.UpdatedAtBlock),
		}
	}

	txs, err = database.GetAllTx(address, startValue, false)

	if err != nil {
		panic(err)
	}

	count := 0
	historicalBalance := make(map[string]big.Float)
	profit := new(big.Float)

	for _, tx := range txs {
		// Buy or Sell token transaction.
		if len(tx.Logs) == 2 {
			sendAmount := utils.GetCoinAmount(tx.Logs[0].Value, tx.Logs[0].ContractDecimals)
			sendUsdPrice := new(big.Float).Mul(&sendAmount, new(big.Float).SetFloat64(tx.Logs[0].UsdPrice))
			fmt.Println("Send token", tx.Logs[0].ContractSymbol, sendAmount.String(), "at", sendUsdPrice)

			receiveAmount := utils.GetCoinAmount(tx.Logs[1].Value, tx.Logs[1].ContractDecimals)
			receiveUsdPrice := new(big.Float).Mul(&receiveAmount, new(big.Float).SetFloat64(tx.Logs[1].UsdPrice))
			fmt.Println("Receive token", tx.Logs[1].ContractSymbol, receiveAmount.String(), "for", receiveUsdPrice)

			lastSendBalance := historicalBalance[tx.Logs[0].ContractSymbol]
			lastSendBalance.Sub(&lastSendBalance, sendUsdPrice)
			historicalBalance[tx.Logs[0].ContractSymbol] = lastSendBalance

			lastReceiveBalance := historicalBalance[tx.Logs[1].ContractSymbol]
			lastReceiveBalance.Add(receiveUsdPrice, &lastReceiveBalance)
			historicalBalance[tx.Logs[1].ContractSymbol] = lastReceiveBalance

			fmt.Println("Last balance of", tx.Logs[0].ContractSymbol, "is", lastSendBalance.String())
			fmt.Println("Last balance of", tx.Logs[1].ContractSymbol, "is", lastReceiveBalance.String())

			// Profit = Sell Price - Price At Nearest Sale.
			// profit := new(big.Float).Add(&lastBalance, receiveUsdPrice)
			// balance := new(big.Float)
		} else if len(tx.Logs) == 1 { // Send to another address
			sendAmount := utils.GetCoinAmount(tx.Logs[0].Value, tx.Logs[0].ContractDecimals)
			sendUsdPrice := new(big.Float).Mul(&sendAmount, new(big.Float).SetFloat64(tx.Logs[0].UsdPrice))
			fmt.Println("Send token", tx.Logs[0].ContractSymbol, sendAmount.String(), "at", sendUsdPrice)

			lastSendBalance := historicalBalance[tx.Logs[0].ContractSymbol]
			lastSendBalance.Sub(sendUsdPrice, &lastSendBalance)
			historicalBalance[tx.Logs[0].ContractSymbol] = lastSendBalance

			fmt.Println("Last balance of", tx.Logs[0].ContractSymbol, "is", lastSendBalance.String())
		}
		count += len(tx.Logs)
	}

	fmt.Println("Total transfer event:", count)
	for k, v := range historicalBalance {
		fmt.Println("Balance of", k, v.String())
		profit.Add(profit, &v)
	}
	fmt.Println("Profit:", profit.String())
}
