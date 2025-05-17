package dbretry

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectWithRetry(maxRetries int, retryInterval time.Duration) *gorm.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", user, pass, host, name)

	var db *gorm.DB
	var err error

	for i := 1; i <= maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to database.")
			return db
		}

		log.Printf("Attempt %d/%d: failed to connect to DB: %v", i, maxRetries, err)
		time.Sleep(retryInterval)
	}

	log.Fatalf("Could not connect to database after %d attempts: %v", maxRetries, err)
	return nil
}
