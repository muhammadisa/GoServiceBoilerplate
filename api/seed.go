package api

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/muhammadisa/go-service-boilerplate/api/models"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/dbconnector"
)

// Seed struct
type Seed struct {
	DB     *gorm.DB
	Tables []interface{}
}

// ISeed interface
type ISeed interface {
	ReinitializeStructs() error
	Migrate() error
	DropTableIfExist() error
}

// StoreTables master table api list
func StoreTables() []interface{} {
	return []interface{}{
		models.Foobar{},
		models.User{},
	}
}

// CheckIsDBNil checking database is nil or not
func CheckIsDBNil(db *gorm.DB) bool {
	if db == nil {
		return true
	}
	return false
}

// ConnectToDB connect to db
func ConnectToDB() *gorm.DB {
	// Loading .env file
	err := godotenv.Load()

	// Checking error for loading .env file
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
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
		DBSqlitePath: "",
	}.Connect()

	if err != nil {
		fmt.Println(err)
	}

	return db
}

// ReinitializeStructs ini all struct tables
func (seed Seed) ReinitializeStructs() error {
	if CheckIsDBNil(seed.DB) {
		seed.DB = ConnectToDB()
	}

	seed.Tables = StoreTables()

	err := seed.DB.DropTableIfExists(seed.Tables...).Error
	if err != nil {
		return err
	}

	err = seed.DB.AutoMigrate(seed.Tables...).Error
	if err != nil {
		return err
	}

	return nil
}

// Migrate migrate tables
func (seed Seed) Migrate() error {
	if CheckIsDBNil(seed.DB) {
		seed.DB = ConnectToDB()
	}

	seed.Tables = StoreTables()
	return seed.DB.AutoMigrate(seed.Tables...).Error
}

// DropTableIfExist drop tables if exist
func (seed Seed) DropTableIfExist() error {
	if CheckIsDBNil(seed.DB) {
		seed.DB = ConnectToDB()
	}

	seed.Tables = StoreTables()
	return seed.DB.DropTableIfExists(seed.Tables...).Error
}
