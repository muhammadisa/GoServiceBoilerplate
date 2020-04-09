package routes

import (
	"fmt"
	"log"
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
)

// Routes struct
type Routes struct {
	Echo  *echo.Echo
	Group *echo.Group
	DB    *gorm.DB
}

// RouteConfigs struct
type RouteConfigs struct {
	EchoData  *echo.Echo
	DB        *gorm.DB
	APISecret string
	Version   string
	Port      string
	Origins   []string
}

// IRouteConfigs interface
type IRouteConfigs interface {
	NewHTTPRoute()
}

// NewHTTPRoute echo route initialization
func (rc RouteConfigs) NewHTTPRoute() {
	// Initialize route configs
	restful := rc.EchoData.Group(fmt.Sprintf("api/%s", rc.Version))
	handler := &Routes{
		Echo:  rc.EchoData,
		Group: restful,
		DB:    rc.DB,
	}
	handler.Echo.Validator = customvalidator.CustomValidator{Validator: validator.New()}
	handler.setupMiddleware(rc.APISecret, rc.Origins)
	handler.setInitRoutes()

	// Internal routers
	handler.initFoobarRoutes()

	// Starting Echo Server
	log.Fatal(handler.Echo.Start(rc.Port))
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
			Message:    "Server is running",
			Data:       true,
		})
	})
}

// Create route initialization function here

func (r *Routes) initFoobarRoutes() {
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(r.DB)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	_foobarApi.NewFoobarHandler(r.Group, foobarUsecase)
}
