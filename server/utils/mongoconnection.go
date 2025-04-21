package utils

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService interface {
	Ping(ctx context.Context) error
	GetCollection(databaseName, collectionName string) (CollectionService, error)
}

type Mongoclient struct {
	client *mongo.Client
}

func (m *Mongoclient) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, nil)
}

type CollectionService interface {
	InsertOne(ctx context.Context, document interface{}) (interface{}, error)
	// Add other methods as needed, e.g., Find, Update, Delete, etc.
}

type MongoCollection struct {
	collection *mongo.Collection
}

func (m *MongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	result, err := m.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (m *Mongoclient) GetCollection(databaseName, collectionName string) (CollectionService, error) {
	collection := m.client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil, errors.New("collection not found")
	}
	return &MongoCollection{collection: collection}, nil
}

var (
	mongoinstance *Mongoclient
	once          sync.Once
)

// InitDb wala kaam  Ye cod
func GetMongoClient() *Mongoclient {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			log.Println("MONGO_URI not set in .env")
		}
		clientOptions := options.Client().ApplyURI(mongoURI)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Printf("Failed to create and connect MongoDB client: %v", err)
		}

		mongoinstance = &Mongoclient{client: client}
		log.Println("MongoDB connection established")
	})

	return mongoinstance
}
