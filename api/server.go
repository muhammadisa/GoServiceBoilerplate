package api

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	_grpc "github.com/muhammadisa/restful-api-boilerplate/api/foobar/delivery/grpc"
	_foobarRepo "github.com/muhammadisa/restful-api-boilerplate/api/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/restful-api-boilerplate/api/foobar/usecase"
	"github.com/muhammadisa/restful-api-boilerplate/api/routes"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	// Load database credential env
	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	apiSecret := os.Getenv("API_SECRET")
	origin := os.Getenv("ORIGINS")
	origins := strings.Split(origin, ",")

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

	// Migrate and checking table fields changes
	err = Seed{DB: db}.Migrate()
	if err != nil {
		log.Fatalf("Unable to migrate %v", err)
		return
	}

	// Foobar
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(db)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)

	mode := os.Getenv("MODE")

	switch mode {
	case "rest":

		// Initialize middleware and route
		e := echo.New()
		routes.NewRoute(e, db, apiSecret, origins, "v2")

		log.Fatal(e.Start(":8080"))

		break

	case "grpc":

		port := ":4040"
		listener, err := net.Listen("tcp", port)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error while listening on %s", port))
		}
		fmt.Println(fmt.Sprintf("gRPC Server is Listening on %s", port))

		server := grpc.NewServer()
		_grpc.NewFoobarServerGrpc(server, foobarUsecase)

		err = server.Serve(listener)
		if err != nil {
			fmt.Println("Unexpected Error", err)
		}

		break

	default:
		panic("Unknown mode on env setting")
	}

}
