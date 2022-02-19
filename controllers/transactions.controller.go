package controllers

import (
	"context"
	"my-assets-be/config"
	"my-assets-be/models"
)

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
