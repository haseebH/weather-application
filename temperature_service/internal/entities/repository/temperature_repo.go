package repository

import (
	"context"
	"time"
)

type Temperature struct {
	ID       string    `json:"id" bson:"_id,omitempty"`
	Location string    `json:"location" bson:"location"`
	Date     time.Time `json:"date" bson:"date"`
	Value    float64   `bson:"value" json:"value"`
}
type TemperatureRepository interface {
	InsertTemperature(ctx context.Context, temp *Temperature) (*Temperature, error)
	GetTemperaturesByPeriod(location string, start, end time.Time) ([]*Temperature, error)
	GetTemperature(location string, start time.Time) (*Temperature, error)
}
