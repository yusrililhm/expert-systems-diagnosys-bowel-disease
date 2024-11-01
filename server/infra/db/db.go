package db

import (
	"log"
	"os"
	"time"
	"usus-sehat/server/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func NewDb() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})

	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}

	if err := sqlDb.Ping(); err != nil {
		return nil, err
	}

	sqlDb.SetConnMaxIdleTime(5 * time.Minute)
	sqlDb.SetConnMaxLifetime(10 * time.Minute)
	sqlDb.SetMaxOpenConns(20)
	sqlDb.SetMaxIdleConns(5)

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Disease{})
	db.AutoMigrate(&entity.Treatment{})
	db.AutoMigrate(&entity.Symptom{})

	admin := &entity.User{
		Username:  os.Getenv("ADMIN_USERNAME"),
		FullName:  os.Getenv("ADMIN_FULLNAME"),
		Phone:     os.Getenv("ADMIN_PHONE"),
		Gender:    true,
		BirthDate: time.Now(),
		Role:      "Admin",
		Password:  os.Getenv("ADMIN_PASSWORD"),
	}

	admin.GenerateFromPassword()

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(admin).Error; err != nil {
		log.Fatalf("[warn] an error occured : %s", err.Error())
		return nil, err
	}

	return db, nil
}
