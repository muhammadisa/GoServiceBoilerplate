package dbconnector

import (
	"fmt"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/mssql"    // MSSql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"    // MySql Driver
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Driver

	"github.com/jinzhu/gorm"
)

// IDatabase interface
type IDatabase interface {
	Connect() (*gorm.DB, error)
}

// DBCredential struct
type DBCredential struct {
	DBDriver     string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSqlitePath string
}

// Connect connect to selected database dialect
func (dbCredential DBCredential) Connect() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	drivers := []string{
		fmt.Sprintf("mysql~%s", fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbCredential.DBUser,
			dbCredential.DBPassword,
			dbCredential.DBHost,
			dbCredential.DBPort,
			dbCredential.DBName,
		)),
		fmt.Sprintf("postgres~%s", fmt.Sprintf(
			"host=%s port=%s user=%s sslmode=disable dbname=%s password=%s",
			dbCredential.DBHost,
			dbCredential.DBPort,
			dbCredential.DBUser,
			dbCredential.DBName,
			dbCredential.DBPassword,
		)),
		fmt.Sprintf("mssql~%s", fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s",
			dbCredential.DBUser,
			dbCredential.DBPassword,
			dbCredential.DBHost,
			dbCredential.DBPort,
			dbCredential.DBName,
		)),
	}

	for index := range drivers {
		drv := strings.Split(drivers[index], "~")
		if dbCredential.DBDriver == drv[0] {
			db, err = gorm.Open(dbCredential.DBDriver, drv[1])
			if err != nil {
				return &gorm.DB{}, err
			}
			break
		}
	}

	fmt.Println("Database Connected")
	return db, nil
}
