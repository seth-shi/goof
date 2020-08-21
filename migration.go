package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path"
	"time"
)




// make migration files
func makeMigration(context *cli.Context) error {

	var err error
	// generate
	table := context.Args().First()

	now := time.Now().Format("20060102150405")

	alias := fmt.Sprintf("%s_%s", table, now)
	filename := fmt.Sprintf("%s/%s_%s.go", env.migrationDir, now, table)

	dir := path.Dir(filename)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(filename, []byte(stub(alias)), os.ModePerm)
	if err == nil {
		log.Info("publish migration to:" + filename)
	}

	return err
}

func stub(alias string) string {

	stub := `package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/seth-shi/goof/runtime"
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