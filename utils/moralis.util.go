package utils

import (
	"encoding/json"
	"fmt"
	"my-assets-be/types"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var client fasthttp.Client = fasthttp.Client{
	NoDefaultUserAgentHeader: true,
	DisablePathNormalizing:   true,
}

func GetTransactions(address string, page int, offset int) {
	a := fiber.AcquireAgent()
	res := fiber.AcquireResponse()
	req := a.Request()
	req.Header.Add("X-API-Key", os.Getenv("MORALIS_WEB3_SECRET"))
	req.Header.Add("accept", "application/json")
	req.SetRequestURI(os.Getenv("MORALIS_HOST") + "/" + address + "?chain=bsc&offset=1")

	if err := client.Do(req, res); err != nil {
		panic(err)
	}

	txsResp := types.AccountTxResp{}
	txs := []types.Transaction{}

	if err := json.Unmarshal(res.Body(), &txsResp); err != nil {
		panic(err)
	}

	for _, v := range txsResp.Result {
		nonce, err := strconv.ParseInt(v.Nonce, 10, 64)
		if err != nil {
			panic(err)
		}
		txs = append(txs, types.Transaction{
			Hash:  v.Hash,
			Nonce: nonce,
		})
	}

	fmt.Println(txs[0].Nonce)
}
