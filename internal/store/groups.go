package store

import (
	"context"
	"errors"
	"paw-me-back/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupStore struct {
	db *gorm.DB
}

func (s *GroupStore) GetAll(ctx context.Context, userID uuid.UUID) ([]model.Group, error) {
	var user model.User

	err := s.db.WithContext(ctx).
		Model(&model.User{}).
		Preload("Groups").
		Preload("Groups.Owner").
		Where("id = ?", userID).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return user.Groups, nil
}

func (s *GroupStore) GetByID(ctx context.Context, groupID uuid.UUID) (*model.Group, error) {
	var group model.Group
	err := s.db.WithContext(ctx).
		Preload("Members").
		Preload("Owner").
		Take(&group, "id = ?", groupID).
		Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &group, nil
}

func (s *GroupStore) IsMember(ctx context.Context, userId uuid.UUID, groupId uuid.UUID) (bool, error) {
	var user model.User

	err := s.db.WithContext(ctx).
		Preload("Groups", "id = ?", groupId).
		Where("id = ?", userId).
		First(&user).
		Error

	if err != nil {
		return false, err
	}

	return len(user.Groups) > 0, nil
}

func (s *GroupStore) AddMembers(ctx context.Context, group *model.Group, userIDs []uuid.UUID) error {
	var users []model.User
	if err := s.db.WithContext(ctx).Find(&users, "id IN ?", userIDs).Error; err != nil {
		return err
	}

	err := s.db.WithContext(ctx).
		Model(&group).
		Association("Members").
		Append(&users)

	return err
}
