package controllers

import (
	"context"
	"my-assets-be/config"
	"my-assets-be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLatestBlockSync(address string) (int64, error) {
	filter := bson.M{"key": "syncStatus"}
	client, err := config.GetMongoClient()

	if err != nil {
		return 0, err
	}

	var status models.TransactionSyncStatus

	collection := client.Database("transactions").Collection(address)

	if err := collection.FindOne(context.TODO(), filter).Decode(&status); err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}

	return status.BlockNumber, nil
}

func UpdateTransactionSyncStatus(address string, blockNumber int64) error {
	filter := bson.M{"key": "syncStatus"}
	client, err := config.GetMongoClient()

	if err != nil {
		return err
	}

	var status models.TransactionSyncStatus

	collection := client.Database("transactions").Collection(address)

	if err := collection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": models.TransactionSyncStatus{
		BlockNumber: blockNumber,
	}}).Decode(&status); err != nil {
		if err == mongo.ErrNoDocuments {
			status.Key = "syncStatus"
			status.BlockNumber = blockNumber
			status.Status = false
			if _, insertErr := collection.InsertOne(context.TODO(), status); insertErr != nil {
				return insertErr
			}
		} else {
			return err
		}
	}

	return nil
}

func CreateNormalTransaction(address string, tx models.NormalTransaction) error {
	client, err := config.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database("transactions").Collection(address)

	_, err = collection.InsertOne(context.TODO(), tx)

	if err != nil {
		return err
	}

	return nil
}
