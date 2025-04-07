package repositories

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"

	"github.com/stretchr/testify/assert"
)

type UserRepositoryInMemory struct {
	users  map[int64]*entities.User
	mu     sync.RWMutex
	nextID int64
}

func NewUserRepositoryInMemory() port.UserRepository {
	return &UserRepositoryInMemory{
		users:  make(map[int64]*entities.User),
		mu:     sync.RWMutex{},
		nextID: 1,
	}
}

func (r *UserRepositoryInMemory) Create(ctx context.Context, user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return nil
}

func (r *UserRepositoryInMemory) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}

func (r *UserRepositoryInMemory) GetByDocument(ctx context.Context, document string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, user := range r.users {
		if user.Document == document {
			return user, nil
		}
	}
	return nil, errors.New("usuário não encontrado com este documento")
}

func (r *UserRepositoryInMemory) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("usuário não encontrado com este email")
}

func (r *UserRepositoryInMemory) ListAll(ctx context.Context) ([]entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	users := make([]entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}

func TestUserRepositoryInMemory_Create(t *testing.T) {
	repo := NewUserRepositoryInMemory()
	ctx := context.Background()

	user := &entities.User{
		FullName:  "John Doe",
		Document:  "123.456.789-00",
		Email:     "john.doe@example.com",
		Password:  "securepassword",
		CreatedAt: time.Now(),
	}

	err := repo.Create(ctx, user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	retrievedUser, err := repo.GetByID(ctx, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.FullName, retrievedUser.FullName)
	assert.Equal(t, user.Document, retrievedUser.Document)
	assert.Equal(t, user.Email, retrievedUser.Email)
}

func TestUserRepositoryInMemory_GetByID_Found(t *testing.T) {
	repo := NewUserRepositoryInMemory()
	ctx := context.Background()

	expectedUser := &entities.User{
		FullName:  "Jane Doe",
		Document:  "987.654.321-00",
		Email:     "jane.doe@example.com",
		Password:  "anotherpassword",
		CreatedAt: time.Now(),
	}
	err := repo.Create(ctx, expectedUser)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetByID(ctx, expectedUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, retrievedUser)
}

func TestUserRepositoryInMemory_GetByID_NotFound(t *testing.T) {
	repo := NewUserRepositoryInMemory()
	ctx := context.Background()

	retrievedUser, err := repo.GetByID(ctx, 999)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "usuário não encontrado")
	assert.Nil(t, retrievedUser)
}

func TestUserRepositoryInMemory_GetByDocument_Found(t *testing.T) {
	repo := NewUserRepositoryInMemory()
	ctx := context.Background()
	document := "111.222.333-44"

	expectedUser := &entities.User{
		FullName:  "Peter Pan",
		Document:  document,
		Email:     "peter.pan@neverland.com",
		Password:  "flyinghigh",
		CreatedAt: time.Now(),
	}
	err := repo.Create(ctx, expectedUser)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetByDocument(ctx, document)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, retrievedUser)
}

func TestUserRepositoryInMemory_GetByDocument_NotFound(t *testing.T) {
	repo := NewUserRepositoryInMemory()
	ctx := context.Background()
	document := "nonexistent"

	retrievedUser, err := repo.GetByDocument(ctx, document)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "usuário não encontrado com este documento")
	assert.Nil(t, retrievedUser)
}

func TestUserRepositoryInMemory_GetByEmail_Found(t *testing.T) {
	repo := NewUserRepositoryInMemory()
	ctx := context.Background()
	email := "alice@wonderland.net"

	expectedUser := &entities.User{
		FullName:  "Alice",
		Document:  "555.444.333-22",
		Email:     email,
		Password:  "downtherabbithole",
		CreatedAt: time.Now(),
	}
	err := repo.Create(ctx, expectedUser)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetByEmail(ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, retrievedUser)
}

func TestUserRepositoryInMemory_GetByEmail_NotFound(t *testing.T) {
	repo := NewUserRepositoryInMemory()
	ctx := context.Background()
	email := "notfound@example.com"

	retrievedUser, err := repo.GetByEmail(ctx, email)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "usuário não encontrado com este email")
	assert.Nil(t, retrievedUser)
}
