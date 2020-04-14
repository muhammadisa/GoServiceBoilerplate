package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammadisa/go-service-boilerplate/api/app/user"
	"github.com/muhammadisa/go-service-boilerplate/api/auth"
	"github.com/muhammadisa/go-service-boilerplate/api/cache"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

type postgreUserRepo struct {
	DB    *gorm.DB
	Cache cache.Redis
}

// NewPostgresUserRepo function
func NewPostgresUserRepo(db *gorm.DB, cacheClient cache.Redis) user.Repository {
	return &postgreUserRepo{
		DB:    db,
		Cache: cacheClient,
	}
}

func (uS *postgreUserRepo) Login(usr *models.User) (*models.User, *auth.Authenticated, error) {
	var err error
	var uSr *models.User = &models.User{}

	err = uS.DB.Model(
		&models.User{},
	).Where(
		"email = ?",
		usr.Email,
	).First(
		&uSr,
	).Error
	if err != nil {
		return nil, nil, err
	}
	return uSr, &auth.Authenticated{}, nil
}

func (uS *postgreUserRepo) Register(usr *models.User) error {
	var err error

	err = uS.DB.Model(
		&models.User{},
	).Create(
		usr,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (uS *postgreUserRepo) Update(usr *models.User) error {
	var err error

	err = uS.DB.Model(
		&models.User{},
	).Where(
		"id = ?",
		usr.ID.String(),
	).Update(models.User{
		Email:     usr.Email,
		Password:  usr.Password,
		UpdatedAt: usr.UpdatedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (uS *postgreUserRepo) Delete(id uuid.UUID) error {
	var err error

	err = uS.DB.Model(
		&models.User{},
	).Where(
		"id = ?",
		id.String(),
	).Delete(
		&models.User{},
	).Error
	if err != nil {
		return err
	}
	return nil
}
