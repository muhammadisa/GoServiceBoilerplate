package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-boilerplate/api/app/user"
	"github.com/muhammadisa/go-service-boilerplate/api/auth"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	"github.com/muhammadisa/go-service-boilerplate/api/response"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/message"
	uuid "github.com/satori/go.uuid"
)

// UserHandler struct
type UserHandler struct {
	uSUsecase user.Usecase
}

// NewUserHandler intialize endpoint
func NewUserHandler(e *echo.Group, uSr user.Usecase) {
	handler := &UserHandler{
		uSUsecase: uSr,
	}
	e.POST("/user/login/", handler.Login)
	e.POST("/user/register/", handler.Register)
}

var (
	model = models.User{}
)

// Login and auth user
func (uS *UserHandler) Login(c echo.Context) error {
	var err error
	var usr models.User

	err = c.Bind(&usr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	err = c.Validate(usr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	authUsr, err := uS.uSUsecase.Login(&usr)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(usr.ID, "POST", model, true),
		Data: auth.Authenticated{
			Data: authUsr,
		},
	})
}

// Register user
func (uS *UserHandler) Register(c echo.Context) error {
	var err error
	var usr models.User

	err = c.Bind(&usr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	err = c.Validate(usr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	err = uS.uSUsecase.Register(&usr)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    message.GenerateMessage(uuid.Nil, "POST", model, false),
			Data:       err,
		})
	}
	return c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    message.GenerateMessage(usr.ID, "POST", model, true),
		Data:       err,
	})
}
