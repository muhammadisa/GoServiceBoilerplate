package foobar

import "github.com/muhammadisa/restful-api-boilerplate/api/models"

// Repository interface
type Repository interface {
	Fetch() (*[]models.Foobar, error)
	GetByID(id uint64) (*models.Foobar, error)
	Update(vT *models.Foobar) error
	Store(vT *models.Foobar) error
	Delete(id uint64) error
}
