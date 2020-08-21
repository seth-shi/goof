package goof

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var orm *gorm.DB
var commands = map[string]func() error{
	"migrate": migrate,
	"migrate:rollback": func() error {

		return nil
	},
}

func Boot(db *gorm.DB) {

	orm = db

	// load migration files
	LoadMigrationFiles("./database/migrations")

	// here register command,
	// Call to the goof
	registerGoofCommand()
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
