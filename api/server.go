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

	// Api Config
	apiSecret := os.Getenv("API_SECRET")
	origin := os.Getenv("ORIGINS")
	origins := strings.Split(origin, ",")

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
	debugEnv := os.Getenv("DEBUG")
	debug, err := strconv.ParseBool(debugEnv)
	if err != nil {
		log.Fatalf("Unable parsing debug env value %v", err)
		return
	}
	db.LogMode(debug)

	// Migrate and checking table fields changes
	err = Seed{DB: db}.Migrate()
	if err != nil {
		log.Fatalf("Unable to migrate %v", err)
		return
	}

	// Checking mode from env
	switch mode := os.Getenv("MODE"); mode {
	case "rest":

		// Init routes
		routes.RouteConfigs{
			EchoData:  echo.New(),
			DB:        db,
			APISecret: apiSecret,
			Version:   "v2",
			Port:      os.Getenv("HTTP_PORT"),
			Origins:   origins,
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
