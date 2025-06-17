package postgres

import (
	"fmt"
	"log"
	"my-go-task/dbModels"

	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *gorm.DB // If we want to access db object via global variable

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func loadYAMLConfig(filename string, key string, target any) {
	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config %s: %v", filename, err)
	}
	if err := v.Sub(key).Unmarshal(target); err != nil {
		log.Fatalf("Error unmarshaling config %s: %v", filename, err)
	}
}

func ConnectDatabase() *gorm.DB {
	var dbConfig DBConfig
	loadYAMLConfig("configs/postgresDb.yaml", "database", &dbConfig)
	log.Println("✅ Database configs loaded from yaml files!")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL with GORM!")

	// DB = db
	db.AutoMigrate(&dbModels.User{}, &dbModels.Watchlist{}, &dbModels.WatchlistScript{})
	log.Println("✅ PostgreSQL database migrated successfully!")

	return db
}
