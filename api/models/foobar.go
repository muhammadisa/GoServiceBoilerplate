package models

import (
	"time"
)

// Foobar struct
type Foobar struct {
	ID            uint64    `gorm:"primary_key;auto_increment" json:"id"`
	FoobarContent string    `gorm:"" json:"foobar_content" validate:"required"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
