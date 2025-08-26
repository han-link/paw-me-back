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

func (s *UserStore) GetBySuperTokenID(ctx context.Context, superTokenID string) (*model.User, error) {
	var user model.User
	err := s.db.WithContext(ctx).
		Where("super_token_id = ?", superTokenID).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) UsernameExists(username string) (bool, error) {
	var count int64
	err := s.db.
		Model(&model.User{}).
		Where("username = ?", username).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *UserStore) Create(user *model.User) error {
	err := s.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) Update(ctx context.Context, user *model.User) error {
	err := s.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
