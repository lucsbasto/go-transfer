package repositories

import (
	"context"

	"go-transfer/internal/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByDocument(ctx context.Context, document string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).Where("document = ?", document).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ListAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
