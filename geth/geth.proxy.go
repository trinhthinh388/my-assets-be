package geth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"my-assets-be/constants/endpoints"
)

type TxReceiptLog struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      int      `json:"blockNumber,string"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	LogIndex         int      `json:"logIndex,string"`
	BlockHash        string   `json:"blockHash"`
	Removed          bool     `json:"removed"`
}

type TxReceipt struct {
	BlockHash         string         `json:"blockHash"`
	BlockNumber       int            `json:"blockNumber,string"`
	ContractAddress   string         `json:"contractAddress"`
	CumulativeGasUsed int            `json:"cumulativeGasUsed,string"`
	GasUsed           int            `json:"gasUsed,string"`
	From              string         `json:"from"`
	Logs              []TxReceiptLog `json:"logs"`
	LogsBloom         string         `json:"logsBloom"`
	Status            string         `json:"status"`
	TransactionIndex  string         `json:"transactionIndex"`
	To                string         `json:"to"`
	TransactionHash   string         `json:"transactionHash"`
	Type              int            `json:"type,string"`
}

type TxReceiptResp struct {
	JsonRPC string    `json:"jsonrpc"`
	Id      int       `json:"id"`
	Result  TxReceipt `json:"result"`
}

func GetTransactionReceipt(txHash string) (TxReceipt, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, reqErr := http.NewRequest("GET", endpoints.BSC_SCAN_HOST, nil)

	if reqErr != nil {
		return TxReceipt{}, reqErr
	}

	q := req.URL.Query()
	q.Add("module", "proxy")
	q.Add("action", "eth_getTransactionReceipt")
	q.Add("txhash", txHash)
	q.Add("apikey", os.Getenv("BSC_SCAN_API_KEY"))
	req.URL.RawQuery = q.Encode()

	resp, respErr := client.Do(req)

	if respErr != nil {
		return TxReceipt{}, respErr
	}

	jsonResp := new(TxReceiptResp)

	json.NewDecoder(resp.Body).Decode(&jsonResp)

	defer resp.Body.Close()

	return jsonResp.Result, nil
}
