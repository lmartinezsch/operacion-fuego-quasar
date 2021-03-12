package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
	"github.com/lmartinezsch/operacion-fuego-quasar/database/models"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbConfig := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(dbDriver, dbConfig)
	db.Exec("USE operacion_fuego_quasar")

	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}

// InitializeOtrs initializes the database
func InitializeOtrs() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_OTRS_HOST")
	dbDriver := os.Getenv("DB_OTRS_DRIVER")
	dbUser := os.Getenv("DB_OTRS_USER")
	dbPassword := os.Getenv("DB_OTRS_PASSWORD")
	dbName := os.Getenv("DB_OTRS_NAME")
	dbPort := os.Getenv("DB_OTRS_PORT")

	dbConfig := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dbConfig)
	db, err := gorm.Open(dbDriver, dbConfig)
	db.Exec("USE otrs")

	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	return db, err
}
