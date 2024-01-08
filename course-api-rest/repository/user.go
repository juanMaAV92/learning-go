package repository

import (
	"context"

	"api.com/go/rest-ws/models"
)

type UserReposisitory interface {
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertUser(ctx context.Context, user *models.User) error
	Close() error
}

var implementation UserReposisitory

func SetUserRepository(repository UserReposisitory) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func Close() error {
	return implementation.Close()
}
