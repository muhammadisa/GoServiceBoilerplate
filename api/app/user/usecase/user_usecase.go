package usecase

import (
	"errors"
	"time"

	"github.com/muhammadisa/go-service-boilerplate/api/app/user"
	"github.com/muhammadisa/go-service-boilerplate/api/auth"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

// userUsecase struct
type userUsecase struct {
	userRepository user.Repository
}

// NewUserUsecase function
func NewUserUsecase(uSr user.Repository) user.Usecase {
	return &userUsecase{
		userRepository: uSr,
	}
}

func (uS userUsecase) Login(usr *models.User) (*models.User, error) {
	uSr, err := uS.userRepository.Login(usr)
	err = auth.VerifyPassword(uSr.Password, usr.Password)
	if err != nil {
		return nil, errors.New("Email or Password is incorrect")
	}
	if err != nil {
		return nil, err
	}
	return uSr, nil
}

func (uS userUsecase) Register(usr *models.User) error {
	err := uS.userRepository.Register(usr)
	if err != nil {
		return err
	}
	return nil
}

func (uS userUsecase) Update(usr *models.User) error {
	usr.UpdatedAt = time.Now()
	err := uS.userRepository.Update(usr)
	if err != nil {
		return err
	}
	return nil
}

func (uS userUsecase) Delete(id uuid.UUID) error {
	return uS.userRepository.Delete(id)
}
