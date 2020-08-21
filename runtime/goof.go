package goof

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var orm *gorm.DB
var commands = map[string]func() error{
	"migrate": migrate,
	"migrate:rollback": migrateRollback,
}

func Boot(db *gorm.DB) error {

	orm = db

	dir := os.Getenv("MIGRATION_DIR")
	if len(dir) == 0 {
		dir = "database/migrations"
	}
	// load migration files
	err := LoadMigrationFiles(dir)
	if err != nil {
		return err
	}

	// here register command,
	// Call to the goof
	registerGoofCommand()

	return nil
}

func registerGoofCommand() {

	for _, arg := range os.Args {

		if handle, exists := commands[arg]; exists {

			err := handle()
			if err != nil {
				log.Fatal(err)
			}

			os.Exit(0)
		}
	}
}
