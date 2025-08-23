package model

import "github.com/google/uuid"

type Expense struct {
	BaseModel
	Amount float64
	Name   string `gorm:"not null"`

	UserID uuid.UUID
	User   *User `gorm:"constraint:OnDelete:CASCADE;"`

	GroupID uuid.UUID
	Group   *Group `gorm:"constraint:OnDelete:CASCADE;"`

	Payees []Payee `gorm:"foreignKey:ExpenseID;references:ID;constraint:OnDelete:CASCADE;"`
}
