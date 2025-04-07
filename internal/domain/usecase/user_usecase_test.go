package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"go-transfer/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) GetByDocument(ctx context.Context, document string) (*entities.User, error) {
	args := m.Called(ctx, document)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) ListAll(ctx context.Context) ([]entities.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.User), args.Error(1)
}

func TestUserUseCase_CreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUser(mockRepo)
	ctx := context.Background()

	input := UserInput{
		FullName: "John Doe",
		Document: "123.456.789-00",
		Email:    "john.doe@example.com",
		Password: "securepassword",
		Type:     entities.CommonWallet,
		Balance:  100.0,
	}

	expectedUser := &entities.User{
		FullName: input.FullName,
		Document: input.Document,
		Email:    input.Email,
		Password: input.Password,
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.User")).Return(nil).Run(func(args mock.Arguments) {
		createdUser := args.Get(1).(*entities.User)
		createdUser.ID = 1
		assert.Equal(t, expectedUser.FullName, createdUser.FullName)
		assert.Equal(t, expectedUser.Document, createdUser.Document)
		assert.Equal(t, expectedUser.Email, createdUser.Email)
		assert.Equal(t, expectedUser.Password, createdUser.Password)
	})

	user, err := userUseCase.CreateUser(ctx, input)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.FullName, user.FullName)
	assert.Equal(t, expectedUser.Document, user.Document)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Password, user.Password)
	assert.NotZero(t, user.ID)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_CreateUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUser(mockRepo)
	ctx := context.Background()

	input := UserInput{
		FullName: "John Doe",
		Document: "123.456.789-00",
		Email:    "john.doe@example.com",
		Password: "securepassword",
		Type:     entities.MerchantWallet,
		Balance:  100.0,
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.User")).Return(errors.New("database error"))

	user, err := userUseCase.CreateUser(ctx, input)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUser(mockRepo)
	ctx := context.Background()
	userID := int64(1)

	expectedUser := &entities.User{
		ID:        userID,
		FullName:  "Jane Doe",
		Document:  "987.654.321-00",
		Email:     "jane.doe@example.com",
		Password:  "anotherpassword",
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetByID", ctx, userID).Return(expectedUser, nil)

	retrievedUser, err := userUseCase.GetUserByID(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, retrievedUser)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetUserByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := NewUser(mockRepo)
	ctx := context.Background()
	userID := int64(1)

	mockRepo.On("GetByID", ctx, userID).Return(nil, errors.New("user not found"))

	retrievedUser, err := userUseCase.GetUserByID(ctx, userID)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Nil(t, retrievedUser)
	mockRepo.AssertExpectations(t)
}
