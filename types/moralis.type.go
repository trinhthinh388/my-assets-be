package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountTxResp struct {
	Total    int      `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
	Result   []TxResp `json:"result"`
}

type TxResp struct {
	Hash             string     `json:"hash" bson:"hash"`
	Nonce            int        `json:"nonce,string" bson:"nonce"`
	TransactionIndex int        `json:"transaction_index,string" bson:"transactionIndex"`
	FromAddress      string     `json:"from_address" bson:"fromAddress"`
	ToAddress        string     `json:"to_address" bson:"toAddress"`
	Value            string     `json:"value" bson:"value"`
	Gas              string     `json:"gas" bson:"gas"`
	GasPrice         string     `json:"gas_price" bson:"gasPrice"`
	Input            string     `json:"input" bson:"input"`
	BlockHash        string     `json:"block_hash" bson:"blockHash"`
	BlockNumber      int        `json:"block_number,string" bson:"blockNumber"`
	BlockTimestamp   *time.Time `json:"block_timestamp" bson:"blockTimestamp"`
}

type TxReceiptLogResp struct {
	LogIndex         int        `json:"log_index,string" bson:"logIndex"`
	Hash             string     `json:"transaction_hash" bson:"transactionHash"`
	TransactionIndex int        `json:"transaction_index,string" bson:"transactionIndex"`
	Address          string     `json:"address" bson:"address"`
	Data             string     `json:"data" bson:"data"`
	Topic0           string     `json:"topic0" bson:"topic0"`
	Topic1           string     `json:"topic1" bson:"topic1"`
	Topic2           string     `json:"topic2" bson:"topic2"`
	Topic3           string     `json:"topic3" bson:"topic3"`
	BlockHash        string     `json:"block_hash" bson:"blockHash"`
	BlockNumber      int        `json:"block_number,string" bson:"blockNumber"`
	BlockTimestamp   *time.Time `json:"block_timestamp" bson:"blockTimestamp"`
}

type TxReceiptResp struct {
	Hash                     string             `json:"hash" bson:"hash"`
	Nonce                    int                `json:"nonce,string" bson:"nonce"`
	TransactionIndex         int                `json:"transaction_index,string" bson:"transactionIndex"`
	FromAddress              string             `json:"from_address" bson:"fromAddress"`
	ToAddress                string             `json:"to_address" bson:"toAddress"`
	Value                    *BigInt            `json:"value" bson:"value"`
	Gas                      string             `json:"gas" bson:"gas"`
	GasPrice                 string             `json:"gas_price" bson:"gasPrice"`
	Input                    string             `json:"input" bson:"input"`
	ReceiptCumulativeGasUsed string             `json:"receipt_cumulative_gas_used" bson:"receiptCumulativeGasUsed"`
	ReceiptGasUsed           string             `json:"receipt_gas_used" bson:"receiptGasUsed"`
	ReceiptContractAddress   string             `json:"receipt_contract_address" bson:"receiptContractAddress"`
	ReceiptRoot              string             `json:"receipt_root" bson:"receiptRoot"`
	ReceiptStatus            string             `json:"receipt_status" bson:"receiptStatus"`
	TransferIndex            []int              `json:"transfer_index" bson:"transferIndex"`
	BlockNumber              int                `json:"block_number,string" bson:"blockNumber"`
	BlockHash                string             `json:"block_hash" bson:"blockHash"`
	Logs                     []TxReceiptLogResp `json:"logs" bson:"logs"`
	BlockTimestamp           *time.Time         `json:"block_timestamp" bson:"blockTimestamp"`
}

type AssetNativePrice struct {
	Value    BigInt `json:"value" bson:"value"`
	Decimals int    `json:"decimals" bson:"decimals"`
	Name     string `json:"name" bson:"name"`
	Symbol   string `json:"symbol" bson:"symbol"`
}
type AssetPrice struct {
	USDPrice    float64          `json:"usdPrice" bson:"usdPrice"`
	NativePrice AssetNativePrice `json:"nativePrice" bson:"nativePrice"`
}

type TxDetail struct {
	ID                       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Hash                     string             `json:"hash" bson:"hash"`
	Nonce                    int                `json:"nonce,string" bson:"nonce"`
	TransactionIndex         int                `json:"transaction_index,string" bson:"transactionIndex"`
	FromAddress              string             `json:"fromAddress" bson:"fromAddress"`
	ToAddress                string             `json:"toAddress" bson:"toAddress"`
	Value                    *BigInt            `json:"value" bson:"value"`
	Gas                      string             `json:"gas" bson:"gas"`
	GasPrice                 string             `json:"gasPrice" bson:"gasPrice"`
	Input                    string             `json:"input" bson:"input"`
	ReceiptCumulativeGasUsed string             `json:"receiptCumulativeGasUsed" bson:"receiptCumulativeGasUsed"`
	ReceiptGasUsed           string             `json:"receiptGasUsed" bson:"receiptGasUsed"`
	ReceiptContractAddress   string             `json:"receiptContractAddress" bson:"receiptContractAddress"`
	ReceiptRoot              string             `json:"receiptRoot" bson:"receiptRoot"`
	ReceiptStatus            string             `json:"receiptStatus" bson:"receiptStatus"`
	TransferIndex            []int              `json:"transferIndex" bson:"transferIndex"`
	BlockNumber              int                `json:"blockNumber,string" bson:"blockNumber"`
	BlockHash                string             `json:"blockHash" bson:"blockHash"`
	Logs                     []TxDetailLog      `json:"logs" bson:"logs"`
	BlockTimestamp           *time.Time         `json:"blockTimestamp" bson:"blockTimestamp"`
}

type TxDetailLog struct {
	ContractAddress  string           `json:"contractAddress" bson:"contractAddress"`
	ContractSymbol   string           `json:"contractSymbol" bson:"contractSymbol"`
	ContractName     string           `json:"contractName" bson:"contractName"`
	ContractIconURL  string           `json:"contractIconUrl" bson:"contractIconUrl"`
	ContractDecimals int              `json:"contractDecimals" bson:"contractDecimals"`
	Value            string           `json:"value" bson:"value"`
	UsdPrice         float64          `json:"usdPrice" bson:"usdPrice"`
	FromAddress      string           `json:"fromAddress" bson:"fromAddress"`
	ToAddress        string           `json:"toAddress" bson:"toAddress"`
	NativePrice      AssetNativePrice `json:"nativePrice" bson:"nativePrice"`
}

type TokenInfo struct {
	Symbol   string `json:"symbol" bson:"symbol"`
	Name     string `json:"name" bson:"string"`
	Logo     string `json:"logo" bosn:"logo"`
	Decimals int    `json:"decimals,string" bson:"decimals"`
}
