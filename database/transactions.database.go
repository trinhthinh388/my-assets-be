package database

import (
	"context"
	"my-assets-be/config"
	"my-assets-be/types"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var addresses map[string]bool
var addressSync sync.Once

func GetTx(address string, startValue types.TxDetail, pageSize int, desc bool) ([]types.TxDetail, error) {
	var opts *options.FindOptions
	client, err := config.GetMongoClient()
	var txs []types.TxDetail
	var filter interface{}

	if err != nil {
		return txs, err
	}

	if !desc {
		opts = options.Find().SetSort(bson.M{"blockNumber": 1}).SetLimit(int64(pageSize))
	} else {
		opts = options.Find().SetSort(bson.M{"blockNumber": -1}).SetLimit(int64(pageSize))
	}

	collection := client.Database("transactions").Collection(address)

	if startValue.BlockNumber == 0 {
		filter = bson.M{}
	} else {
		if !desc {
			filter = bson.M{"blockNumber": bson.M{
				"$gt": startValue.BlockNumber,
			}}
		} else {
			filter = bson.M{"blockNumber": bson.M{
				"$lt": startValue.BlockNumber,
			}}
		}
	}

	cur, err := collection.Find(context.TODO(), filter, opts)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return txs, nil
		}
		return txs, err
	}

	for cur.Next(context.TODO()) {
		var tx types.TxDetail
		if err := cur.Decode(&tx); err != nil {
			return txs, err
		}

		txs = append(txs, tx)
	}

	if err := cur.Err(); err != nil {
		return txs, err
	}

	cur.Close(context.TODO())

	return txs, nil
}

func GetAllTx(address string, startValue types.TxDetail, desc bool) ([]types.TxDetail, error) {
	var opts *options.FindOptions
	if desc {
		opts = options.Find().SetSort(bson.M{"blockNumber": -1})
	} else {
		opts = options.Find().SetSort(bson.M{"blockNumber": 1})
	}

	client, err := config.GetMongoClient()
	var txs []types.TxDetail
	var filter interface{}

	if err != nil {
		return txs, err
	}

	if startValue.BlockNumber == 0 {
		filter = bson.M{}
	} else {
		if !desc {
			filter = bson.M{"blockNumber": bson.M{
				"$gt": startValue.BlockNumber,
			}}
		} else {
			filter = bson.M{"blockNumber": bson.M{
				"$lt": startValue.BlockNumber,
			}}
		}
	}

	collection := client.Database("transactions").Collection(address)

	cur, err := collection.Find(context.TODO(), filter, opts)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return txs, nil
		}
		return txs, err
	}

	for cur.Next(context.TODO()) {
		var tx types.TxDetail
		if err := cur.Decode(&tx); err != nil {
			return txs, err
		}

		txs = append(txs, tx)
	}

	if err := cur.Err(); err != nil {
		return txs, err
	}

	cur.Close(context.TODO())

	return txs, nil
}

func SyncAddress() {
	if addresses == nil {
		addresses = make(map[string]bool)
	}
	client, err := config.GetMongoClient()

	if err != nil {
		panic(err)
	}

	adds, err := client.Database("transactions").ListCollectionNames(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	for _, v := range adds {
		addresses[v] = true
	}
}

func IsAddressExist(address string) bool {
	addresses := GetAddressMap()

	return addresses[address]
}

func GetAddressMap() map[string]bool {
	addressSync.Do(SyncAddress)

	return addresses
}
