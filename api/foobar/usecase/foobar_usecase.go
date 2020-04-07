package usecase

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/restful-api-boilerplate/api/foobar"
	"github.com/muhammadisa/restful-api-boilerplate/api/models"
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

func (fB foobarUsecase) GetByID(id uint64) (*models.Foobar, error) {
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

func (fB foobarUsecase) Delete(id uint64) error {
	return fB.foobarRepository.Delete(id)
}
