package store

import (
	"context"
	"errors"
	"paw-me-back/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("resource not found")
	/*ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5*/
)

type Groups interface {
	GetAll(ctx context.Context, userID uuid.UUID) ([]model.Group, error)
	GetByID(ctx context.Context, groupID uuid.UUID) (*model.Group, error)
	AddMembers(ctx context.Context, group *model.Group, userIDs []uuid.UUID) error
	IsMember(ctx context.Context, userId uuid.UUID, groupId uuid.UUID) (bool, error)
	Create(ctx context.Context, group *model.Group) error
}

type Users interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetBySuperTokenID(ctx context.Context, superTokenID string) (*model.User, error)
	UsernameExists(username string) (bool, error)
	Create(user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

type Storage struct {
	Groups Groups
	Users  Users
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Groups: &GroupStore{db},
		Users:  &UserStore{db},
	}
}
