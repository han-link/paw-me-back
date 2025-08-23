package model

import "github.com/google/uuid"

type Payee struct {
	BaseModel

	UserID uuid.UUID `gorm:"type:uuid;index;not null"`
	User   *User     `gorm:"constraint:OnDelete:CASCADE;"`

	ExpenseID uuid.UUID `gorm:"type:uuid;index;not null"`
	Expense   *Expense  `gorm:"constraint:OnDelete:CASCADE;"`

	Share float64
}
