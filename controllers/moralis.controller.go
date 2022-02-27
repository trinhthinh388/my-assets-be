package controllers

import (
	"encoding/json"
	"fmt"

	"my-assets-be/config"
	"my-assets-be/constants/moralis"
	"my-assets-be/models"
	"my-assets-be/types"
	"net/url"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var client fasthttp.Client = fasthttp.Client{
	NoDefaultUserAgentHeader: true,
	DisablePathNormalizing:   true,
}

func GetERC20Price(contractAddress string, chain string, toBlock int) (types.AssetPrice, error) {
	var priceResp types.AssetPrice
	a := fiber.AcquireAgent()
	res := fiber.AcquireResponse()
	req := a.Request()
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_WEB3_SECRET"))
	req.Header.Add("accept", "application/json")
	req.SetRequestURI(os.Getenv("MORALIS_HOST") + "/erc20/" + contractAddress + "/price")

	values := url.Values{}
	values.Add("chain", chain)
	if toBlock > 0 {
		values.Add("from_block", strconv.Itoa(toBlock))
	}
	if toBlock > 0 {
		values.Add("to_block", strconv.Itoa(toBlock))
	}

	a.QueryString(values.Encode())

	fmt.Println(a.Request().URI().String())

	if err := client.Do(req, res); err != nil {
		return priceResp, err
	}

	if err := json.Unmarshal(res.Body(), &priceResp); err != nil {
		return priceResp, err
	}

	return priceResp, nil
}

func GetTransactions(address string, offset int, fromBlock int, toBlock int, chain string) (types.AccountTxResp, error) {
	var txsResp types.AccountTxResp
	a := fiber.AcquireAgent()
	res := fiber.AcquireResponse()
	req := a.Request()
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_WEB3_SECRET"))
	req.Header.Add("accept", "application/json")
	req.SetRequestURI(os.Getenv("MORALIS_HOST") + "/" + address)

	values := url.Values{}
	values.Add("chain", chain)
	values.Add("limit", strconv.Itoa(moralis.PAGE_SIZE))
	values.Add("offset", strconv.Itoa(offset))

	if fromBlock > 0 {
		values.Add("from_block", strconv.Itoa(fromBlock))
	}
	if toBlock > 0 {
		values.Add("to_block", strconv.Itoa(toBlock))
	}

	a.QueryString(values.Encode())

	fmt.Println(a.Request().URI().String())

	if err := client.Do(req, res); err != nil {
		return txsResp, err
	}

	if err := json.Unmarshal(res.Body(), &txsResp); err != nil {
		return txsResp, err
	}

	return txsResp, nil
}

func GetTxDetail(address string, chain string) (types.TxReceiptResp, error) {
	var txResp types.TxReceiptResp
	a := fiber.AcquireAgent()
	res := fiber.AcquireResponse()
	req := a.Request()
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_WEB3_SECRET"))
	req.Header.Add("accept", "application/json")
	req.SetRequestURI(os.Getenv("MORALIS_HOST") + "/transaction/" + address)

	values := url.Values{}
	values.Add("chain", chain)
	a.QueryString(values.Encode())

	fmt.Println(a.Request().URI().String())

	if err := client.Do(req, res); err != nil {
		return txResp, err
	}

	if err := json.Unmarshal(res.Body(), &txResp); err != nil {
		return txResp, err
	}

	return txResp, nil
}

func SyncTx(address string, chain string) {
	dbClient := config.GetMysqlClient()
	var latestTx models.Transaction
	var offset int = 0
	var fromBlock int = 0

	// Get the latest block that has been synced.
	dbClient.Limit(1).Order("block_number desc, timestamp desc").Where("owner = ? AND chain = ?", address, chain).Find(&latestTx)

	if latestTx.BlockNumber <= 0 {
		fromBlock = 0
	} else {
		fromBlock = latestTx.BlockNumber
	}

	for {
		var txsDetail []models.Transaction
		resp, err := GetTransactions(address, offset, fromBlock, 0, chain)
		if err != nil {
			panic(err)
		}

		validTxs := resp.Result

		if fromBlock > 0 {
			validTxs = resp.Result[:len(resp.Result)-1]
		}

		if len(validTxs) <= 0 {
			break
		}

		for _, v := range validTxs {
			txsDetail = append(txsDetail, models.Transaction{
				Hash:        v.Hash,
				From:        v.FromAddress,
				To:          v.ToAddress,
				Value:       v.Value,
				Gas:         v.Gas,
				GasPrice:    v.GasPrice,
				Input:       v.Input,
				BlockNumber: v.BlockNumber,
				Timestamp:   v.BlockTimestamp,
				Chain:       chain,
				Owner:       address,
			})
		}

		dbClient.Create(txsDetail)

		SendTextMessage(address, "NEW TRANSACTIONS FOUNDED")

		if offset += moralis.PAGE_SIZE; offset > resp.Total {
			break
		}
	}

	SendTextMessage(address, "SYNCED")
}

func GetTokenInfo(contractAddress string, chain string) (*types.TokenInfo, error) {
	var tokenResp []types.TokenInfo
	a := fiber.AcquireAgent()
	res := fiber.AcquireResponse()
	req := a.Request()
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_WEB3_SECRET"))
	req.Header.Add("accept", "application/json")
	req.SetRequestURI(os.Getenv("MORALIS_HOST") + "/erc20/metadata")

	values := url.Values{}
	values.Add("addresses", contractAddress)
	values.Add("chain", chain)
	a.QueryString(values.Encode())

	fmt.Println(a.Request().URI().String())

	if err := client.Do(req, res); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res.Body(), &tokenResp); err != nil {
		return nil, err
	}

	if len(tokenResp) <= 0 {
		return &types.TokenInfo{}, nil
	}

	return &tokenResp[0], nil
}
