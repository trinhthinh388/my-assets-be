package config

import (
	"fmt"
	"my-assets-be/models"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var instance *gorm.DB
var instanceError error

func GetMysqlClient() *gorm.DB {
	once.Do(func() {
		username := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_HOST")
		port := os.Getenv("MYSQL_PORT")
		database := os.Getenv("MYSQL_DATABASE")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

		instance, instanceError = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if instanceError != nil {
			panic(instanceError)
		}
		fmt.Printf("Connect to %s success.\n", database)
		Migrate(instance)
	})

	return instance
}

func Migrate(client *gorm.DB) {
	fmt.Println("Migrating database...")

	err := client.AutoMigrate(models.Transaction{}, models.Contract{})

	if err != nil {
		fmt.Printf("Error migrating database : error=%v", err)
		return
	}

	fmt.Println("Your database have been updated. ✅")
}
