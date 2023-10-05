package main

import (
	"log"

	"BACKEND/config"
	"BACKEND/internals/pkg/database"
	"BACKEND/internals/pkg/database/migration"

	"github.com/go-gormigrate/gormigrate/v2"
)

func setUp() (*gormigrate.Gormigrate, error) {
	// configuration
	config.Setup()

	// database
	db, err := database.OpenConnection()
	if err != nil {
		return nil, err
	}

	// migration
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migration.InitTables,
	})

	return m, nil
}

func main() {
	m, err := setUp()
	if err != nil {
		log.Fatalf("Error while initializing migrator: %v", err.Error())
	}
	if err = m.Migrate(); err != nil {
		log.Fatalf("Error while migrate: %v", err.Error())
	}
	log.Printf("Migration completed")
}
