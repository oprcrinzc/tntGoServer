package db

import (
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Conn() *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MongoDbString")))
	if err != nil {
		panic(err)
	}
	return client
}
