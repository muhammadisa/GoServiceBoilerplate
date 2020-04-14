package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Foobar struct
type Foobar struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	FoobarContent string    `gorm:"" json:"foobar_content" validate:"required"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeCreate generate uuid v4
func (fooBar Foobar) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
