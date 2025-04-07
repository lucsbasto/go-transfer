package port

import (
	"context"

	"go-transfer/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id int64) (*entities.User, error)
	GetByDocument(ctx context.Context, document string) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
}
