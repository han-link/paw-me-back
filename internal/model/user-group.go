package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGroup struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	GroupID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserGroup) TableName() string {
	return "user_groups"
}
