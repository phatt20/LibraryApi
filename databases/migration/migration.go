package main

import (
	"github.com/phatt20/LibraryApi/config"
	"github.com/phatt20/LibraryApi/databases"
	"github.com/phatt20/LibraryApi/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(&conf.Database)

	tx := db.Connect().Begin()

	customerMigration(tx)
	adminMigration(tx)
	bookMigration(tx)
	customerCoinMigration(tx)
	borrowRecordMigration(tx)
	purchaseHistoryMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func customerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Customer{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func bookMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Book{})
}

func customerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.CustomerCoin{})
}

func borrowRecordMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.BorrowRecord{})
}

func purchaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
