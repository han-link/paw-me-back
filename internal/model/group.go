package model

import "github.com/google/uuid"

type Group struct {
	BaseModel
	Name     string
	OwnerID  uuid.UUID
	Owner    User      `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:RESTRICT;"`
	Members  []User    `gorm:"many2many:user_groups;"`
	Expenses []Expense `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:CASCADE;"`
}
