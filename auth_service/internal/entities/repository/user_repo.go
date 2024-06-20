package repository

import (
	"context"
)

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password,omitempty" bson:"password" binding:"required"`
	Token    string `json:"token" bson:"-"`
	Location string `json:"location" bson:"location" binding:"required"`
}

//go:generate mockery --name UserRepository --inpackage --filename=user_repo_mock.go
type UserRepository interface {
	RegisterUser(ctx context.Context, user *User) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	LoginUser(ctx context.Context, username, password string) (*User, error)
	ValidateToken(ctx context.Context, token string) error
}
