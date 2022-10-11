package database

import (
	"errors"
	"log"
	"os"

	//"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go_rest_monolith/config"
)

var (
	Gorm2 *gorm.DB
)

// connectDb
func ConnectDb() {
	if config.ENVIRONMENT == "development" {
		ConnectSQlite()
	} else if config.ENVIRONMENT == "production" {
		ConnectMysql()
	} else {
		ConnectSQlite()
	}
}

// connect mysql
func ConnectMysql() {
	/* dsn := "root:mysql@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	} else {
		log.Println("Mysql connected")
		Gorm2 = db
	} */
}

func ConnectSQlite() {
	_, err := os.Stat("config/database/sqlite/test_db.db")

	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("config/database/sqlite/test_db.db")

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
	}

	db, err := gorm.Open(sqlite.Open("config/database/sqlite/test_db.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	} else {
		log.Println("SQlite connected")
		Gorm2 = db
	}
}
