package routes

import (
	"net/http"

	_foobarApi "github.com/muhammadisa/restful-api-boilerplate/api/foobar/delivery/http"
	_foobarRepo "github.com/muhammadisa/restful-api-boilerplate/api/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/restful-api-boilerplate/api/foobar/usecase"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammadisa/restful-api-boilerplate/api/response"
	"github.com/muhammadisa/restful-api-boilerplate/api/utils/customvalidator"
	"github.com/muhammadisa/restful-api-boilerplate/api/utils/message"
)

// Routes struct
type Routes struct {
	Echo  *echo.Echo
	Group *echo.Group
	DB    *gorm.DB
}

// NewRoute echo route initialization
func NewRoute(
	echoData *echo.Echo,
	db *gorm.DB,
	apiSecret string,
	origins []string,
	masterVersion string,
) {
	restful := echoData.Group("/api/v2")
	handler := &Routes{
		Echo:  echoData,
		Group: restful,
		DB:    db,
	}
	handler.Echo.Validator = customvalidator.CustomValidator{Validator: validator.New()}
	handler.setupMiddleware(apiSecret, origins)
}

func (r *Routes) setupMiddleware(apiSecret string, origins []string) {
	// main middleware
	r.Echo.Use(middleware.Recover())
	r.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
}

func (r *Routes) setInitRoutes() {
	r.Echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.Response{
			StatusCode: http.StatusOK,
			Message:    message.GenerateMessage(0, "GET", "starting point", true),
			Data:       "Running",
		})
	})
}

// Create route initialization function here

func (r *Routes) initFoobarRoutes() {
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(r.DB)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	_foobarApi.NewFoobarHandler(r.Group, foobarUsecase)
}
