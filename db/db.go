package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const EmployeesCollection = "employees"
const DBName = "test"
const MongoURI = "mongodb+srv://rishi:7qNuRoBoaN8n8cg9@cluster0.46eeubf.mongodb.net/" +
	DBName +
	"?retryWrites=true&w=majority"

var Instance *mongo.Database

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		return fmt.Errorf("unable to create mongo client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("unable to connect mongo server: %w", err)
	}

	db := client.Database(DBName)
	// Instance = types.MongoInstance{
	// 	Client: client,
	// 	DB:     db,
	// }
	Instance = db

	return nil
}
