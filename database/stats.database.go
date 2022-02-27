package database

import (
	"context"
	"my-assets-be/config"
	"my-assets-be/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetStat(address string, chain string) (*types.AddressStats, error) {
	client, err := config.GetMongoClient()
	filter := bson.M{"address": address}
	var stats types.AddressStats

	if err != nil {
		return &stats, err
	}

	collection := client.Database("stats").Collection(chain)

	if err := collection.FindOne(context.TODO(), filter).Decode(&stats); err != nil {
		if err == mongo.ErrNoDocuments {
			return &stats, nil
		}
		return &stats, err
	}

	return &stats, nil
}

func InsertStats(address string, chain string, stats types.AddressStats) error {
	client, err := config.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database("stats").Collection(chain)

	if _, err := collection.InsertOne(context.TODO(), stats); err != nil {
		return err
	}

	return nil
}

func UpdateStats(id primitive.ObjectID, chain string, stats types.AddressStats) error {
	client, err := config.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database("stats").Collection(chain)

	if _, err := collection.UpdateByID(context.TODO(), id, stats); err != nil {
		return err
	}

	return nil
}
