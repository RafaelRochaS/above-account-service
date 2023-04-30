package repository

import (
	"errors"
	"log"

	"github.com/RafaelRochaS/above-account-service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbConn struct {
	db *gorm.DB
}

func (dbConn *DbConn) InstantiateDbConn() error {
	var err error
	dbConn.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Printf("DbRepository :: Error: Failed to connect database")
		return err
	}

	dbConn.Migrate()

	return nil
}

func (dbConn DbConn) Migrate() {
	dbConn.db.AutoMigrate(&models.User{})
}

func (dbConn DbConn) Create(user *models.User) error {
	return dbConn.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})
}

func (dbConn DbConn) CheckConn() error {
	if sqlDb, err := dbConn.db.DB(); err == nil {
		if err = sqlDb.Ping(); err == nil {
			return nil
		}
	}

	return errors.New("failed to connect to DB")
}
