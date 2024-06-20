package datastore

import (
	"context"

	"github.com/haseebh/weatherapp_auth/internal/entities/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	collection *mongo.Collection
}

func NewUserDB(client *Database, collectionName string) repository.UserRepository {
	return &userRepo{
		collection: client.db.Collection(collectionName),
	}
}

func (ur *userRepo) RegisterUser(ctx context.Context, user *repository.User) (*repository.User, error) {
	_, err := ur.collection.InsertOne(ctx, user)
	return user, err
}
func (ur *userRepo) FindUserByEmail(ctx context.Context, email string) (*repository.User, error) {
	var user repository.User
	err := ur.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *userRepo) LoginUser(ctx context.Context, email, password string) (*repository.User, error) {
	var user repository.User
	err := ur.collection.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (ur *userRepo) ValidateToken(ctx context.Context, token string) error {

	return nil
}
