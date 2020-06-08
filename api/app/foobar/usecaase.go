package foobar

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

// Usecase interface
type Usecase interface {
	Fetch() (*gorm.DB, *[]models.Foobar, error)
	GetByID(id uuid.UUID) (*models.Foobar, error)
	Update(foobar *models.Foobar) error
	Store(foobar *models.Foobar) error
	Delete(id uuid.UUID) error
}
