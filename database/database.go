package database

import (
	"go-chain/config"
	"go-chain/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(cfg *config.Config) error {
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return err
	}

	// Auto-migrate tables
	err = DB.AutoMigrate(
		&models.Category{},
		&models.Recipe{},
		&models.Favorite{},
		&models.ShoppingItem{},
	)
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
