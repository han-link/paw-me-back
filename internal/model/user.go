package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	SuperTokenID uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	Username     string
	Email        string
	Groups       []Group   `gorm:"many2many:user_groups;"`
	Expenses     []Expense `json:"expenses,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Payees       []Payee   `json:"payees,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}
