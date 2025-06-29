package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(host, port, dbname string) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port + "/" + dbname)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	return client.Database(dbname)
}
