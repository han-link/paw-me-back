package model

type User struct {
	BaseModel
	SuperTokenID string `gorm:"uniqueIndex;"`
	Username     string
	Email        string

	Groups   []Group   `gorm:"many2many:user_groups;"`
	Expenses []Expense `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Payees   []Payee   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}
