package lib

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database modal
type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewDatabase creates a new database instance
func NewDatabase(env Env, logger Logger) Database {
	dbURI := env.MongoDBURI
	dbName := env.MongoDBName
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		logger.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Database connection established")

	return Database{Client: client, Database: client.Database(dbName)}
}

func (db *Database) Collection(name string) *mongo.Collection {
	return db.Database.Collection(name)
}
