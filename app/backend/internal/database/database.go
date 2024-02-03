package database

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

/*
func NewDatabase() (*gorm.DB, error) {
	logrus.Info("Setting up Connection with the Database")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	logrus.Info(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
*/

func NewDatabase() (*gorm.DB, error) {

	logrus.Info("Setting up Connection with the Database")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	logrus.Info(connectionString)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
