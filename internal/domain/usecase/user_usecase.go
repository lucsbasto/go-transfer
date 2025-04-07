package usecase

import (
	"context"

	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"
)

type UserInput struct {
	FullName string              `json:"full_name"`
	Document string              `json:"document"`
	Email    string              `json:"email"`
	Password string              `json:"password"`
	Type     entities.WalletType `json:"type"`
	Balance  float64             `json:"balance"`
}

type User struct {
	userRepo port.UserRepository
}

func NewUser(userRepo port.UserRepository) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) CreateUser(ctx context.Context, input UserInput) (*entities.User, error) {
	user := &entities.User{
		FullName: input.FullName,
		Document: input.Document,
		Email:    input.Email,
		Password: input.Password,
	}

	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetUserByID(ctx context.Context, id int64) (*entities.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
