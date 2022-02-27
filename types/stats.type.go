package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddressStats struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Address        string             `json:"address" bson:"address"`
	BuyAmount      BigInt             `json:"buyAmount" bson:"buyAmount"`
	SellAmount     BigInt             `json:"sellAmount" bson:"sellAmount"`
	UpdatedAtBlock int                `json:"updatedAtBlock" bson:"updatedAtBlock"`
	UpdatedAtTx    string             `json:"updatedAtTx" bson:"updatedAtTx"`
}
