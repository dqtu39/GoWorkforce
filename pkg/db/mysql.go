package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func NewMySQLConn() (*sql.DB, error) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"),
	//	os.Getenv("DB_NAME"))
	//fmt.Println(dsn)

	dsn := "goworkforce_user:goworkforce_password@tcp(mysql:3306)/goworkforce?parseTime=true"
	fmt.Println(dsn)

	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			if pingErr := db.Ping(); pingErr == nil {
				log.Println("Connected to MySQL")
				return db, nil
			}
		}
		log.Println("MySQL not ready, retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("could not connect to MySQL: %v", err)
}
