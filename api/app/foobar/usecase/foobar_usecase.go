package usecase

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-boilerplate/api/app/foobar"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

// foobarUsecase struct
type foobarUsecase struct {
	foobarRepository foobar.Repository
}

// NewFoobarUsecase function
func NewFoobarUsecase(fbUsecase foobar.Repository) foobar.Usecase {
	return &foobarUsecase{
		foobarRepository: fbUsecase,
	}
}

func (fbUsecase foobarUsecase) Fetch() (*gorm.DB, *[]models.Foobar, error) {
	db, res, err := fbUsecase.foobarRepository.Fetch()
	if err != nil {
		return nil, nil, err
	}
	return db, res, nil
}

func (fbUsecase foobarUsecase) GetByID(id uuid.UUID) (*models.Foobar, error) {
	res, err := fbUsecase.foobarRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (fbUsecase foobarUsecase) Store(fbUsecasear *models.Foobar) error {
	err := fbUsecase.foobarRepository.Store(fbUsecasear)
	if err != nil {
		return err
	}
	return nil
}

func (fbUsecase foobarUsecase) Update(fbUsecasear *models.Foobar) error {
	fbUsecasear.UpdatedAt = time.Now()
	return fbUsecase.foobarRepository.Update(fbUsecasear)
}

func (fbUsecase foobarUsecase) Delete(id uuid.UUID) error {
	return fbUsecase.foobarRepository.Delete(id)
}
