package database

import (
	"context"
	"fmt"
	"github.com/namcchan/go-chatting/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var MongoClient *mongo.Client

func CreateMongoConnection() (*mongo.Client, error) {
	env := configs.GetEnv()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := fmt.Sprintf("%s://%s:%s@%s:%s", env.DBConnection, env.DBUser, env.DBPass, env.DBHost, env.DBPort)

	if env.DBUser == "" || env.DBPass == "" {
		mongodbURI = fmt.Sprintf("%s://%s:%s", env.DBConnection, env.DBHost, env.DBPort)
	}

	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	MongoClient = client
	return client, nil
}

func GetMongoClient() *mongo.Client {
	return MongoClient
}

// GetCollection getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	collection := GetMongoClient().Database("chatting").Collection(collectionName)
	return collection
}

func CloseMongoConnection(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
