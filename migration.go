package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"time"
)




// make migration files
func makeMigration(context *cli.Context) error {

	// generate
	table := context.Args().First()

	now := time.Now().Format("20060102150405")

	alias := fmt.Sprintf("%s_%s", table, now)
	filename := fmt.Sprintf("database/migrations/%s_%s.go", now, table)

	err := ioutil.WriteFile(filename, []byte(stub(alias)), os.ModePerm)
	if err == nil {
		log.Info("publish migration to:" + filename)
	}

	return err
}

func stub(alias string) string {

	stub := `package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/seth-shi/goof/goof"
)

func init()  {

	goof.RegisterMigrate(up_%s, down_%s)
}

// Run the migrations.
func up_%s(db *gorm.DB)  {
	
}

// Reverse the migrations.
func down_%s(db *gorm.DB) {
	
}
`
	return fmt.Sprintf(stub, alias, alias, alias, alias)
}