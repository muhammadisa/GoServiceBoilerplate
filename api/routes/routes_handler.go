package routes

import (
	"fmt"
	"log"
	"net/http"

	_foobarApi "github.com/muhammadisa/go-service-boilerplate/api/app/foobar/delivery/http"
	_foobarRepo "github.com/muhammadisa/go-service-boilerplate/api/app/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/go-service-boilerplate/api/app/foobar/usecase"
	"github.com/muhammadisa/go-service-boilerplate/api/cache"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muhammadisa/go-service-boilerplate/api/response"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/customvalidator"
)

// Routes struct
type Routes struct {
	Echo  *echo.Echo
	Group *echo.Group
	DB    *gorm.DB
	Cache cache.Redis
}

// RouteConfigs struct
type RouteConfigs struct {
	EchoData  *echo.Echo
	DB        *gorm.DB
	APISecret string
	Version   string
	Port      string
	Origins   []string
	Cache     cache.Redis
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
		Cache: rc.Cache,
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
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(r.DB, r.Cache)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	_foobarApi.NewFoobarHandler(r.Group, foobarUsecase)
}
