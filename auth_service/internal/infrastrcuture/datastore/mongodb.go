package datastore

import (
	"context"
	"fmt"
	"github.com/haseebh/weatherapp_auth/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Database struct {
	db *mongo.Database
}

func NewDatabase(ctx context.Context, config *config.Config) (db *Database) {
	opts := options.
		Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d/%s", config.Database.Host, config.Database.Port, config.Database.DBName)).
		SetHeartbeatInterval(time.Second).
		SetMinPoolSize(config.Database.MinPoolSize).
		SetMaxPoolSize(config.Database.MaxPoolSize)
	if config.Database.User != "" && config.Database.Password != "" {
		opts = opts.SetAuth(options.Credential{
			Username: config.Database.User,
			Password: config.Database.Password,
		})
	}
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	mDB := client.Database(config.Database.DBName)

	return &Database{
		db: mDB,
	}
}
