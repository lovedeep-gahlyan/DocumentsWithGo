package server

import (
	"docs/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(config *viper.Viper) *gorm.DB {
	connectionString := config.GetString("database.connection_string")
	maxIdleConnections := config.GetInt("database.max_idle_connections")
	maxOpenConnections := config.GetInt("database.max_open_connections")
	connectionMaxLifetime := config.GetDuration("database.connection_max_lifetime")

	if connectionString == "" {
		log.Fatalf("Database connection string is missing")
	}

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while initializing database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error while migrating database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error while getting database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetConnMaxLifetime(connectionMaxLifetime)

	err = sqlDB.Ping()
	if err != nil {
		sqlDB.Close()
		log.Fatalf("Error while validating database: %v", err)
	}

	return db
}
