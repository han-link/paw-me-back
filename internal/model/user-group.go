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

	User  User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Group Group `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

func (UserGroup) TableName() string { return "user_groups" }
