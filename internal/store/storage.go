package store

import (
	"context"
	"errors"
	"paw-me-back/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Groups interface {
		GetAll(ctx context.Context, userID uuid.UUID) ([]model.Group, error)
		GetByID(ctx context.Context, groupID uuid.UUID) (*model.Group, error)
		AddMembers(ctx context.Context, group *model.Group, userIDs []uuid.UUID) error
		IsMember(ctx context.Context, userId uuid.UUID, groupId uuid.UUID) (bool, error)
	}
	Users interface {
		GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
		GetOrCreateBySuperTokenID(
			ctx context.Context,
			superTokenID uuid.UUID,
			username string,
			email string,
		) (*model.User, error)
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Groups: &GroupStore{db},
		Users:  &UserStore{db},
	}
}
