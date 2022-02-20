package types

type AccountTxResp struct {
	Total    int               `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
	Result   []TransactionResp `json:"result"`
}

type TransactionResp struct {
	Hash             string `json:"hash" bson:"hash"`
	Nonce            string `json:"nonce" bson:"nonce"`
	TransactionIndex string `json:"transactionIndex" bson:"transactionIndex"`
	FromAddress      string `json:"fromAddress" bson:"fromAddress"`
	ToAddress        string `json:"toAddress" bson:"toAddress"`
	Value            string `json:"value" bson:"value"`
	Gas              string `json:"gas" bson:"gas"`
	GasPrice         string `json:"gasPrice" bson:"gasPrice"`
	Input            string `json:"input" bson:"input"`
	BlockHash        string `json:"blockHash" bson:"blockHash"`
	BlockNumber      string `json:"blockNumber" bson:"blockNumber"`
	BlockTimestamp   string `json:"timestamp" bson:"timestamp"`
}
