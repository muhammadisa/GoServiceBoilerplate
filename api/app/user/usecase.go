package user

import (
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

// Usecase interface
type Usecase interface {
	Login(usr *models.User) (*models.User, error)
	Register(usr *models.User) error
	Update(usr *models.User) error
	Delete(id uuid.UUID) error
}
