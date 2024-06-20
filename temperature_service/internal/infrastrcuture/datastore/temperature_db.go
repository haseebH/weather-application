package datastore

import (
	"context"
	"time"

	"github.com/haseebh/weatherapp_temperature/internal/entities/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type temperatureRepo struct {
	collection *mongo.Collection
}

func NewTemperatureDB(client *Database, collectionName string) repository.TemperatureRepository {
	return &temperatureRepo{
		collection: client.db.Collection(collectionName),
	}
}

func (tr *temperatureRepo) InsertTemperature(ctx context.Context, temp *repository.Temperature) (*repository.Temperature, error) {
	_, err := tr.collection.InsertOne(ctx, temp)
	return temp, err
}
func (tr *temperatureRepo) GetTemperaturesByPeriod(location string, start, end time.Time) ([]*repository.Temperature, error) {

	filter := bson.M{
		"location": location,
		"date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}
	cursor, err := tr.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var temperatures []*repository.Temperature
	for cursor.Next(context.Background()) {
		var temp repository.Temperature
		if err := cursor.Decode(&temp); err != nil {
			return nil, err
		}
		temperatures = append(temperatures, &temp)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return temperatures, nil
}

func (tr *temperatureRepo) GetTemperature(location string, date time.Time) (*repository.Temperature, error) {
	filter := bson.M{
		"location": location,
		"date":     date,
	}
	var temperature repository.Temperature
	err := tr.collection.FindOne(context.Background(), filter).Decode(&temperature)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &temperature, nil
}
