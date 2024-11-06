package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func NewMySQLConn() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		getEnvOrDefault("DB_USER", "goworkforce_user"),
		getEnvOrDefault("DB_PASSWORD", "goworkforce_password"),
		getEnvOrDefault("DB_HOST", "mysql"),
		getEnvOrDefault("DB_PORT", "3306"),
		getEnvOrDefault("DB_NAME", "goworkforce"))

	var db *sql.DB
	var err error
	
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			if err := db.Ping(); err == nil {
				db.SetMaxOpenConns(25)
				db.SetMaxIdleConns(25)
				db.SetConnMaxLifetime(5 * time.Minute)
				log.Println("Connected to MySQL")
				return db, nil
			}
		}
		log.Printf("MySQL connection attempt %d failed: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to MySQL after 10 attempts: %v", err)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
