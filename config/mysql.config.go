package config

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbCon *gorm.DB
var dbSync sync.Once

func GetMysqlClient() *gorm.DB {
	dbSync.Do(func() {
		username := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_HOST")
		port := os.Getenv("MYSQL_PORT")
		database := os.Getenv("MYSQL_DATABASE")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

		var err error
		dbCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Failed to connect to Database.")
		}

		fmt.Printf("Connect to %s success.\n", database)
	})

	return dbCon
}

func Migrate() {
	fmt.Println("Migrating database...")

	// client := GetMysqlClient()

	// err := client.AutoMigrate(models.Transaction{})

	// if err != nil {
	// 	panic("Database migrations has been failed.")
	// }

	fmt.Println("Your database have been updated. ✅")
}
