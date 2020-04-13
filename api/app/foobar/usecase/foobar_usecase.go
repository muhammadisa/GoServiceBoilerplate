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
func NewFoobarUsecase(fB foobar.Repository) foobar.Usecase {
	return &foobarUsecase{
		foobarRepository: fB,
	}
}

func (fB foobarUsecase) Fetch() (*gorm.DB, *[]models.Foobar, error) {
	db, res, err := fB.foobarRepository.Fetch()
	if err != nil {
		return nil, nil, err
	}
	return db, res, nil
}

func (fB foobarUsecase) GetByID(id uuid.UUID) (*models.Foobar, error) {
	res, err := fB.foobarRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (fB foobarUsecase) Store(FBar *models.Foobar) error {
	err := fB.foobarRepository.Store(FBar)
	if err != nil {
		return err
	}
	return nil
}

func (fB foobarUsecase) Update(FBar *models.Foobar) error {
	FBar.UpdatedAt = time.Now()
	return fB.foobarRepository.Update(FBar)
}

func (fB foobarUsecase) Delete(id uuid.UUID) error {
	return fB.foobarRepository.Delete(id)
}
