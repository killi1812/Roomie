package Helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	connString := GetConfig().DbConnString

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("").Collection("")
	title := "test"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(result)
	if err == mongo.ErrNoDocuments {
		return
	}

	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}

// GetConnect returns a connnection and function for closing that connection
func GetConnect() (*mongo.Client, func()) {

	connString := GetConfig().DbConnString
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	return client, func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}
}
