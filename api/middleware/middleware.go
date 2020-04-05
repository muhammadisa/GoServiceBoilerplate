package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/restful-api-boilerplate/api/response"
)

// GoMiddleware struct
type GoMiddleware struct{}

// CORS avoid CORS error
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// APISecretKeyCheck checking secret key from public access
func (m *GoMiddleware) APISecretKeyCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		secretKey := c.Request().Header.Get("Secret-Key")

		// Loading .env file
		err := godotenv.Load()

		// Checking error for loading .env file
		if err != nil {
			log.Fatalf("Error getting env, not coming through %v", err)
			c.JSON(http.StatusInternalServerError, response.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal Server Error",
				Data:       nil,
			})
			return nil
		}

		// Checking API Secret Key
		apiSecretKey := os.Getenv("API_SECRET")
		if secretKey != apiSecretKey {
			c.JSON(http.StatusUnauthorized, response.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			})
			return nil
		}
		return next(c)
	}
}

// InitMiddleware initialize middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
