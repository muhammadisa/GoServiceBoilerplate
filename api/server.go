package api

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/restful-api-boilerplate/api/routes"
	"github.com/muhammadisa/restful-api-boilerplate/api/utils/dbconnector"
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

	// Load database credential env and use it
	db, err := dbconnector.DBCredential{
		DBDriver:     os.Getenv("DB_DRIVER"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBSqlitePath: "",
	}.Connect()
	if err != nil {
		fmt.Println(err)
	}

	// Load debuging mode env
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	db.LogMode(debug)

	// Migrate and checking table fields changes
	Seed{DB: db}.Migrate()

	// Checking mode from env
	switch mode := os.Getenv("MODE"); mode {
	case "rest":

		// Init routes
		routes.RouteConfigs{
			EchoData:  echo.New(),
			DB:        db,
			APISecret: os.Getenv("API_SECRET"),
			Version:   "v2",
			Port:      os.Getenv("HTTP_PORT"),
			Origins:   strings.Split(os.Getenv("ORIGINS"), ","),
		}.NewHTTPRoute()
		break

	case "grpc":

		routes.GRPCConfigs{
			DB:       db,
			Protocol: "tcp",
			Port:     os.Getenv("GRPC_PORT"),
		}.NewGRPC()
		break

	default:
		panic("Unknown mode on env setting")
	}

}
