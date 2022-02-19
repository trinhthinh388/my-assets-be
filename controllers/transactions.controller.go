package controllers

import (
	"context"
	"log"
	"my-assets-be/config"
	"my-assets-be/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func CreateNormalTransactions(address string, txs []models.NormalTransaction) error {

	insertableList := make([]interface{}, len(txs))

	for i, v := range txs {
		insertableList[i] = v
	}

	client, err := config.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database("transactions").Collection(address)

	_, err = collection.InsertMany(context.TODO(), insertableList)

	if err != nil {
		return err
	}

	return nil
}

func GetNormalTransactionCount(address string) (int64, error) {
	client, err := config.GetMongoClient()

	if err != nil {
		return 0, err
	}

	collection := client.Database("transactions").Collection(address)

	count, err := collection.CountDocuments(context.TODO(), bson.D{{}})

	return count, err
}

func GetLatestNormalTransaction(address string) (models.NormalTransaction, error) {

	tx := models.NormalTransaction{}

	opts := options.FindOne().SetSort(bson.M{"$natural": -1})

	client, err := config.GetMongoClient()

	if err != nil {
		return tx, err
	}

	collection := client.Database("transactions").Collection(address)

	if err = collection.FindOne(context.TODO(), bson.M{}, opts).Decode(&tx); err != nil {
		log.Fatal(err)
	}

	return tx, nil

}

func GetAllNormalTransaction(address string) ([]models.NormalTransaction, error) {

	filter := bson.D{{}}

	txs := []models.NormalTransaction{}

	client, err := config.GetMongoClient()

	if err != nil {
		return txs, err
	}

	collection := client.Database("transaction").Collection(address)

	cur, findError := collection.Find(context.TODO(), filter)

	if findError != nil {
		return txs, findError
	}

	for cur.Next(context.TODO()) {
		t := models.NormalTransaction{}
		err := cur.Decode(&t)
		if err != nil {
			return txs, err
		}
		txs = append(txs, t)
	}

	cur.Close(context.TODO())
	if len(txs) == 0 {
		return txs, mongo.ErrNoDocuments
	}
	return txs, nil
}
