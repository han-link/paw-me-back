package store

import (
	"context"
	"errors"
	"paw-me-back/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func (s *UserStore) GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	var user model.User

	err := s.db.WithContext(ctx).
		Where("id = ?", userID).
		First(&user).
		Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStore) GetOrCreateBySuperTokenID(
	ctx context.Context,
	superTokenID uuid.UUID,
	username string,
	email string,
) (*model.User, error) {
	var user model.User
	err := s.db.WithContext(ctx).
		Where("super_token_id = ?", superTokenID).
		First(&user).
		Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			user.SuperTokenID = superTokenID
			user.Username = username
			user.Email = email
			err = s.db.Create(&user).Error
			if err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}

	return &user, nil
}
