package database

import (
	"fmt"
	"go-gorm-mux/src/api/config"
	"go-gorm-mux/src/api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *config.Config) {
	var err error

	if config.DB.DbDialect == "mysql" {
		// DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
		// DB, err = gorm.Open(config.DB.DbDialect, DBURL)
		// if err != nil {
		// 	fmt.Printf("Cannot connect to %s database", config.DB.DbDialect)
		// 	log.Fatal("This is the error:", err)
		// } else {
		// 	log.Printf("We are connected to the %s database", config.DB.DbDialect)
		// }
	}

	if config.DB.DbDialect == "postgres" {
		// DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)
		// DB, err = gorm.Open(config.DB.DbDialect, DBURL)
		// if err != nil {
		// 	fmt.Printf("Cannot connect to %s database", config.DB.DbDialect)
		// 	log.Fatal("This is the error:", err)
		// } else {
		// 	log.Printf("We are connected to the %s database", config.DB.DbDialect)
		// }
	}

	if config.DB.DbDialect == "sqlite3" {
		DB, err = gorm.Open(sqlite.Open(config.DB.DbName), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", config.DB.DbDialect)
			log.Fatal("This is the error:", err)
		} else {
			log.Printf("We are connected to the %s database\n", config.DB.DbDialect)
		}

		if res := DB.Exec("PRAGMA foreign_keys = ON"); res.Error != nil {
			fmt.Println(res.Error)
			log.Fatal(res.Error)
		}
	}
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
	log.Println("Database Migration Completed!")
}