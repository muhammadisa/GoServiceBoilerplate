package foobar

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/restful-api-boilerplate/api/models"
)

// Usecase interface
type Usecase interface {
	Fetch() (*gorm.DB, *[]models.Foobar, error)
	GetByID(id uint64) (*models.Foobar, error)
	Update(vT *models.Foobar) error
	Store(vT *models.Foobar) error
	Delete(id uint64) error
}
