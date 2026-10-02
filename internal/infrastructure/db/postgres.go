package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgres() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "5430"),
			getEnv("DB_USER", "postgres"),
			getEnv("DB_PASSWORD", "1111"),
			getEnv("DB_NAME", "lab2"),
		)
	}

	log.Printf("🔌 DSN: %s", dsn)

	gormLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLog})
	if err != nil {
		return nil, err
	}

	// if err := gormDB.AutoMigrate(
	// 	&models.User{},
	// 	&models.Component{},
	// 	&models.Like{},
	// ); err != nil {
	// 	return nil, err
	// }

	sqlDB, _ := gormDB.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return gormDB, nil
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
