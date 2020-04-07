package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_foobarApi "github.com/muhammadisa/restful-api-boilerplate/api/foobar/delivery/http"
	_foobarRepo "github.com/muhammadisa/restful-api-boilerplate/api/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/restful-api-boilerplate/api/foobar/usecase"
	"github.com/muhammadisa/restful-api-boilerplate/api/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_middleware "github.com/muhammadisa/restful-api-boilerplate/api/middleware"
	"github.com/muhammadisa/restful-api-boilerplate/api/response"
	"github.com/muhammadisa/restful-api-boilerplate/api/utils/dbconnector"
	"github.com/muhammadisa/restful-api-boilerplate/api/utils/message"
)

// Run used for start connecting to selected database
func Run() {

	// Loading .env file
	err := godotenv.Load()

	// Checking error for loading .env file
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
		return
	}

	// Load database credential env
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := dbconnector.DBCredential{
		DBDriver:     dbDriver,
		DBHost:       dbHost,
		DBPort:       dbPort,
		DBUser:       dbUser,
		DBPassword:   dbPass,
		DBName:       dbName,
		DBPathSqlite: "",
	}.Connect()
	if err != nil {
		fmt.Println(err)
	}

	// Load debuging mode env
	debugEnv := os.Getenv("DEBUG")
	debug, err := strconv.ParseBool(debugEnv)
	if err != nil {
		log.Fatalf("Unable parsing debug env value %v", err)
		return
	}
	db.LogMode(debug)

	db.AutoMigrate(
		models.Foobar{},
	)

	// Initialize middleware and route
	e := echo.New()
	middL := _middleware.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.Response{
			StatusCode: http.StatusOK,
			Message:    message.GenerateMessage(0, "GET", "home", true),
			Data:       "Running",
		})
	})

	// Foobar
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(db)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	_foobarApi.NewFoobarHandler(e, foobarUsecase)

	// Start echo web framework
	log.Fatal(e.Start(":8080"))

}
