package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoClient represents a MongoDB client

// NewMongoClient creates and returns a new MongoDB client
func NewMongoClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to confirm connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB successfully")
	// database := client.Database(dbName)

	return client, nil
}

// Close disconnects from MongoDB
func Close(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}
	log.Println("Disconnected from MongoDB")
	return nil
}

// GetCollection returns a handle to the specified collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("database").Collection(collectionName)
}

// GetClient returns the underlying MongoDB client
func GetClient(client *mongo.Client) *mongo.Client {
	return client
}

// GetDatabase returns the MongoDB database
func GetDatabase(client *mongo.Client, dbName string) *mongo.Database {
	return client.Database(dbName)
}
