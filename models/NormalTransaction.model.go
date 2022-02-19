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
