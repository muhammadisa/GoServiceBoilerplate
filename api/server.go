package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
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

	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	_, err = dbconnector.DBCredential{
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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "200",
			"message": message.Message{
				IsSuccess:       true,
				HTTPMethod:      "GET",
				TargetModelName: "home page",
				WithID:          0,
			}.GenerateMessage(),
		})
	})

	// Start echo web framework
	log.Fatal(e.Start(":8080"))

}
