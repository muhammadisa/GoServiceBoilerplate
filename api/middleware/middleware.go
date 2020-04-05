package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
			c.JSON(http.StatusInternalServerError, map[string]string{
				"status": "Internal Server Error",
			})
			return nil
		}

		// Checking API Secret Key
		apiSecretKey := os.Getenv("API_SECRET")
		if secretKey != apiSecretKey {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"status": "Unauthorized",
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
