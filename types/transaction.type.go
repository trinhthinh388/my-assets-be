package types

import (
	"math/big"
	"time"
)

type Time time.Time

type BigInt big.Int

type BlockTransaction struct {
	Hash             string `json:"hash" bson:"hash"`
	Nonce            string `json:"nonce" bson:"nonce"`
	BlockHash        string `json:"blockHash" bson:"blockHash"`
	BlockNumber      string `json:"blockNumber" bson:"blockNumber"`
	From             string `json:"from" bson:"from"`
	To               string `json:"to" bson:"to"`
	Value            string `json:"value" bson:"value"`
	Gas              string `json:"gas" bson:"gas"`
	GasPrice         string `json:"gasPrice" bson:"gasPrice"`
	Input            string `json:"input" bson:"input"`
	TransactionIndex string `json:"transactionIndex" bson:"transactionIndex"`
	Type             string `json:"type" bson:"type"`
	R                string `json:"r" bson:"r"`
	S                string `json:"s" bson:"s"`
	V                string `json:"v" bson:"v"`
}

type BlockInfo struct {
	Number           string             `json:"number" bson:"number"`
	Transactions     []BlockTransaction `json:"transactions" bson:"transactions"`
	Hash             string             `json:"hash" bson:"hash"`
	ParentHash       string             `json:"parentHash" bson:"parentHash"`
	Nonce            string             `json:"nonce" bson:"nonce,string"`
	SHA3Uncles       string             `json:"sha3Uncles" bson:"sh3Uncles"`
	LogsBloom        string             `json:"logBlooms" bson:"logBlooms"`
	TransactionsRoot string             `json:"transactionsRoot" bson:"transactionsRoot"`
	StateRoot        string             `json:"stateRoot" bson:"stateRoot"`
	Timestamp        string             `json:"timestamp" bson:"timestamp"`
	ReceiptsRoot     string             `json:"receiptsRoot" bson:"receiptsRoot"`
	Miner            string             `json:"miner" bson:"miner"`
	Difficulty       string             `json:"difficulty" bson:"difficulty"`
	TotalDifficulty  string             `json:"totalDifficulty" bson:"totalDifficulty"`
	ExtraData        string             `json:"extraData" bson:"extraData"`
	Size             string             `json:"size" bson:"size"`
	GasLimit         string             `json:"gasLimit" bson:"gasLimit"`
	GasUsed          string             `json:"gasUsed" bson:"gasUsed"`
	Uncles           []string           `json:"Uncles" bson:"uncles"`
}
