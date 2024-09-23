package config

import (
	"fmt"
	models "zonaquant/models/dbmodels"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewConnection(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return db, nil
}

func InitDB() {
	db, err := NewConnection("testquant.db")
	if err != nil {
		fmt.Printf("error opening sqlite:" + err.Error())
		return
	}

	DB = db
}

func Migrate(db *gorm.DB) error {
	db.Exec("PRAGMA foreign_keys = ON")

	// Auto-migrate all models
	return db.AutoMigrate(
		&models.AccountType{},
		&models.Account{},
		&models.Role{},
		&models.User{},
		&models.UsersAccounts{},
		&models.Producer{},
		&models.Empacadora{},
		&models.PurchaseReport{},
		&models.PurchaseReportDetail{},
		&models.SellingReport{},
		&models.Size{},
		&models.Class{},
		&models.SellingReportDetail{},
		&models.DeliveryReport{},
		&models.Batch{},
		&models.Logistics{},
		&models.TransactionState{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.Advance{},
	)
}
