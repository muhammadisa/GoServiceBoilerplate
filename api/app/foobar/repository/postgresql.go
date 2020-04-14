package repository

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-boilerplate/api/app/foobar"
	"github.com/muhammadisa/go-service-boilerplate/api/cache"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

type postgreFoobarRepo struct {
	DB    *gorm.DB
	Cache cache.Redis
}

// NewPostgresFoobarRepo function
func NewPostgresFoobarRepo(db *gorm.DB, cacheClient cache.Redis) foobar.Repository {
	return &postgreFoobarRepo{
		DB:    db,
		Cache: cacheClient,
	}
}

func (pFb *postgreFoobarRepo) Fetch() (*gorm.DB, *[]models.Foobar, error) {
	var err error
	var fBars *[]models.Foobar = &[]models.Foobar{}

	db := pFb.DB.Model(
		&models.Foobar{},
	).Order(
		"created_at asc",
	).Find(
		&fBars,
	)
	err = db.Error
	if err != nil {
		return nil, nil, err
	}
	return db, fBars, nil
}

func (pFb *postgreFoobarRepo) GetByID(id uuid.UUID) (*models.Foobar, error) {
	var err error
	var fBar *models.Foobar = &models.Foobar{}

	cache := pFb.Cache.Get(cache.Key(models.Foobar{}, id))
	if cache != "nil" && cache != "" {
		json.Unmarshal([]byte(cache), &fBar)
		return fBar, nil
	}
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
	pFb.Cache.Set(*fBar)
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
		FBar.ID.String(),
	).Update(models.Foobar{
		FoobarContent: FBar.FoobarContent,
		UpdatedAt:     FBar.UpdatedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (pFb *postgreFoobarRepo) Delete(id uuid.UUID) error {
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
