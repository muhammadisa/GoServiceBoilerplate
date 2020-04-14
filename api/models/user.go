package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-boilerplate/api/auth"
	uuid "github.com/satori/go.uuid"
)

// User struct
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Email     string    `gorm:"size:255;not null;unique" json:"email" validate:"required,email"`
	Password  string    `gorm:"size:255;not null" json:"password" validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeCreate generate uuid v4 and hashing password
func (user User) BeforeCreate(scope *gorm.Scope) (err error) {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	scope.SetColumn("Password", string(user.Password))
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
