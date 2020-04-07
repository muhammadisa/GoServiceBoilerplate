package models

import (
	"html"
	"strings"
	"time"
)

// Foobar struct
type Foobar struct {
	ID            uint64    `gorm:"primary_key;auto_increment" json:"id" validate:"required,numeric"`
	FoobarContent string    `gorm:"default:''" json:"foobar_content" validate:"required"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare foobar struct
func (foobar *Foobar) Prepare() {
	foobar.ID = 0
	foobar.FoobarContent = html.EscapeString(strings.TrimSpace(foobar.FoobarContent))
	foobar.CreatedAt = time.Now()
	foobar.UpdatedAt = time.Now()
}
