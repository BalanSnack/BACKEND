package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDSN() string {
	// DSN Format: "username:password@tcp(host:port)/dbname"
	return fmt.Sprintf(
		"%s?charset=%s&parseTime=%t&loc=%s&timeout=%d",
		"balansnack:balansnack@tcp(localhost:3306)/balansnack",
		"utf8mb4",
		true,
		"Local",
		0,
	)
}

func newDialector() gorm.Dialector {
	return mysql.Open(getDSN())
}

func OpenConnection() (*gorm.DB, error) {
	db, err := gorm.Open(newDialector())
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	return db, nil
}
