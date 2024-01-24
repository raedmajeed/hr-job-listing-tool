package config

import (
	"fmt"
	dom "github.com/raedmajeed/hr-job-tool/pkg/dom/hr_models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnect(cfg *ConfigParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connection to mysql %s failed, error: %s", cfg.DBName, err)
	}

	err = database.AutoMigrate(dom.Job{}, dom.Profile{})

	if err != nil {
		log.Fatalf("unable to migrate db, err: %v", err)
	}

	return database, nil
}
