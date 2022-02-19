package models

import (
	"math/big"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Time time.Time

type BigInt big.Int

type NormalTransaction struct {
	Id                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BlockNumber       int                `json:"blockNumber,string" bson:"blockNumber,string"`
	TimeStamp         Time               `json:"timeStamp" bson:"timeStamp"`
	Hash              string             `json:"hash" bson:"hash"`
	Nonce             int                `json:"nonce,string" bson:"nonce,string"`
	BlockHash         string             `json:"blockHash" bson:"blockHash"`
	TransactionIndex  int                `json:"transactionIndex,string" bson:"transactionIndex,string"`
	From              string             `json:"from" bson:"from"`
	To                string             `json:"to" bson:"to"`
	Value             *BigInt            `json:"value" bson:"value"`
	Gas               int                `json:"gas,string" bson:"gas,string"`
	GasPrice          *BigInt            `json:"gasPrice" bson:"gasPrice"`
	IsError           int                `json:"isError,string" bson:"isError,string"`
	TxReceiptStatus   string             `json:"txreceipt_status" bson:"txreceipt_status"`
	Input             string             `json:"input" bson:"input"`
	ContractAddress   string             `json:"contractAddress" bson:"contractAddress"`
	CumulativeGasUsed int                `json:"cumulativeGasUsed,string" bson:"cumulativeGasUsed,string"`
	GasUsed           int                `json:"gasUsed,string" bson:"gasUsed,string"`
	Confirmations     int                `json:"confirmations,string" bson:"confirmations,string"`
}
