package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/restful-api-boilerplate/api/foobar"
	"github.com/muhammadisa/restful-api-boilerplate/api/models"
)

type postgreFoobarRepo struct {
	DB *gorm.DB
}

// NewPostgresFoobarRepo function
func NewPostgresFoobarRepo(db *gorm.DB) foobar.Repository {
	return &postgreFoobarRepo{
		DB: db,
	}
}

func (pFb *postgreFoobarRepo) Fetch() (*[]models.Foobar, error) {
	var err error
	var fBars *[]models.Foobar = &[]models.Foobar{}

	err = pFb.DB.Model(
		&models.Foobar{},
	).Find(
		&fBars,
	).Error
	if err != nil {
		return nil, err
	}

	return fBars, nil
}

func (pFb *postgreFoobarRepo) GetByID(id uint64) (*models.Foobar, error) {
	var err error
	var fBar *models.Foobar = &models.Foobar{}

	err = pFb.DB.Model(
		&models.Foobar{},
	).Where(
		"id = ?",
		id,
	).First(
		&fBar,
	).Error
	if err != nil {
		return nil, err
	}

	return fBar, nil
}

func (pFb *postgreFoobarRepo) Store(FBar *models.Foobar) error {
	var err error

	err = pFb.DB.Model(
		&models.Foobar{},
	).Create(
		FBar,
	).Error
	if err != nil {
		return err
	}

	return nil
}

func (pFb *postgreFoobarRepo) Update(FBar *models.Foobar) error {
	var err error

	err = pFb.DB.Model(
		&models.Foobar{},
	).Where(
		"id = ?",
		FBar.ID,
	).Update(models.Foobar{
		FoobarContent: FBar.FoobarContent,
		UpdatedAt:     FBar.UpdatedAt,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (pFb *postgreFoobarRepo) Delete(id uint64) error {
	var err error

	err = pFb.DB.Model(
		&models.Foobar{},
	).Where(
		"id = ?",
		id,
	).Delete(
		&models.Foobar{},
	).Error
	if err != nil {
		return err
	}

	return nil
}
