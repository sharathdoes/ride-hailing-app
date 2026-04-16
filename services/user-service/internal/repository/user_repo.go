package repository

import (
	"context"
	"rideshare/user-service/internal/domain"
)


type UserRepo interface {
    CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
    FindByPhone(ctx context.Context, phone string) (*domain.User, error)
    FindByID(ctx context.Context, id string) (*domain.User, error)
    UpdateUser(ctx context.Context, id string, updates *domain.User) (*domain.User, error)
}

