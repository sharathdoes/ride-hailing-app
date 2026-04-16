package repository

import (
	"context"
	"rideshare/user-service/internal/domain"
)


type DriverRepo interface {
    CreateDriver(ctx context.Context, driver *domain.Driver) (*domain.Driver, error)
    FindByUserID(ctx context.Context, userID string) (*domain.Driver, error)
    UpdateDriver(ctx context.Context, userID string, updates *domain.Driver) (*domain.Driver, error)
    UpdateAvailability(ctx context.Context, userID string, available bool) error
}