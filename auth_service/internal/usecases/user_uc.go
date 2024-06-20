package usecases

import (
	"context"
	"errors"
	"github.com/haseebh/weatherapp_auth/internal/entities/repository"
	"github.com/haseebh/weatherapp_auth/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name UserUseCase --inpackage --filename=user_uc_mock.go

type UserUseCase interface {
	Register(ctx context.Context, user *repository.User) (*repository.User, error)
	Login(ctx context.Context, email, password string) (*repository.User, error)
	ValidateToken(ctx context.Context, token string) error
}

type userUC struct {
	userRepo     repository.UserRepository
	messageQueue repository.MessageQueue
}

func NewUserUseCase(
	repo repository.UserRepository,
	mq repository.MessageQueue,
) UserUseCase {

	return &userUC{
		userRepo:     repo,
		messageQueue: mq,
	}
}

func (uc *userUC) Register(ctx context.Context, user *repository.User) (*repository.User, error) {
	// validation
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email and password are required")
	}

	//Check if user already exists
	existingUser, err := uc.userRepo.FindUserByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	_, err = uc.userRepo.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}
	err = uc.messageQueue.Publish(user)
	return user, err
}

func (uc *userUC) Login(ctx context.Context, email, password string) (*repository.User, error) {
	user, err := uc.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		return nil, err
	}
	user.Token = token
	user.Password = ""
	user.ID = ""
	return user, nil
}
func (uc *userUC) ValidateToken(ctx context.Context, token string) error {
	_, err := middleware.VerifyToken(token)
	if err != nil {
		return err
	}
	return nil
}
