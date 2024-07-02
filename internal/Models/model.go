package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid; primary_key; not null; index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BankModel struct {
	BankID    uuid.UUID `json:"bank_id" gorm:"type:uuid; primary_key; not null; index"`
	BankCode  int64     `json:"bank_code" gorm:"not null; index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignUpModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateNameModel struct {
	FirstName  *string `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName   *string `json:"last_name"`
}

type User struct {
	Base
	FirstName           *string `json:"first_name"`
	MiddleName          *string `json:"middle_name"`
	LastName            *string `json:"last_name"`
	Email               string  `json:"email" gorm:"unique; index"`
	AuthenticationToken string  `json:"authentication_token" gorm:"unique; index"`
	RefreshToken        string  `json:"refresh_token" gorm:"unique; index"`
	SessionKey          string  `json:"session_key"`
	Password            string  `json:"password"`
}

type Bank struct {
	BankModel
	Name string `json:"name" gorm:"unique; index"`
}

func (BankModel *BankModel) BeforeCreate(scope *gorm.DB) error {
	uuid, error := uuid.NewV7()
	if error != nil {
		log.Fatal("Can't create UUID")
	}
	scope.Statement.SetColumn("bank_id", uuid, true)
	return error
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.DB) error {
	uuid, error := uuid.NewV7()
	if error != nil {
		log.Fatal("Can't create UUID")
	}
	scope.Statement.SetColumn("id", uuid, true)
	return error
}
