package dbconnector

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mssql"    // MSSql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"    // MySql Driver
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // Sqlite Driver

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
	DBPathSqlite string
}

// Connect connect to selected database dialect
func (dbCredential DBCredential) Connect() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var connectionString string

	switch dbCredential.DBDriver {

	case "mysql":
		connectionString = fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbCredential.DBUser,
			dbCredential.DBPassword,
			dbCredential.DBHost,
			dbCredential.DBPort,
			dbCredential.DBName,
		)
		db, err = gorm.Open(dbCredential.DBDriver, connectionString)
		break

	case "postgres":
		connectionString = fmt.Sprintf(
			"host=%s port=%s user=%s sslmode=disable dbname=%s password=%s",
			dbCredential.DBHost,
			dbCredential.DBPort,
			dbCredential.DBUser,
			dbCredential.DBName,
			dbCredential.DBPassword,
		)
		db, err = gorm.Open(dbCredential.DBDriver, connectionString)
		break

	case "sqlite":
		connectionString = fmt.Sprintf(
			"%s",
			dbCredential.DBPathSqlite,
		)
		db, err = gorm.Open(dbCredential.DBDriver, connectionString)
		break

	case "mssql":
		connectionString = fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s",
			dbCredential.DBUser,
			dbCredential.DBPassword,
			dbCredential.DBHost,
			dbCredential.DBPort,
			dbCredential.DBName,
		)
		db, err = gorm.Open(dbCredential.DBDriver, connectionString)
		break

	default:
		db, err = &gorm.DB{}, fmt.Errorf("%s DB Driver not available", dbCredential.DBDriver)

	}
	if err != nil {
		return &gorm.DB{}, err
	}

	fmt.Println("Database Connected")
	return db, nil
}
