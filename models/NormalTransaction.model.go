package models

import (
	"math/big"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Time time.Time

type BigInt big.Int

type NormalTransaction struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Hash        string             `json:"hash" bson:"hash,omitempty"`
	BlockNumber int64              `json:"blockNumber,string" bson:"blockNumber,string"`
}

type TransactionSyncStatus struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key         string             `json:"key" bson:"key,omitempty"`
	Status      bool               `json:"status" bson:"status"`
	BlockNumber int64              `json:"blockNumber,string" bson:"blockNumber,string"`
}
