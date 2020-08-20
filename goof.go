//go:generate go-bindata -prefix framework -o framework.go framework/...

package main

import (
	"fmt"
	"github.com/seth-shi/goof/utils"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	log = utils.GetLogInstance()
	workDir string
)

func init()  {

	wd, err := os.Getwd()
	if err != nil {
		log.ErrorFatal("cannot get work dir")
	}

	workDir = wd
}

func main() {

	app := &cli.App{
		Name: "goof",
		Version: "v1.0.1",
		Usage: "development of project foundation scaffolding",
		UsageText: "goof [action] [arguments...]",
		Action: func(c *cli.Context) error {

			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			{
				Name: "init",
				Aliases: []string{"i"},
				Usage: "init project foundation scaffolding",
				UsageText: "goof init [-output]",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "output", Aliases: []string{"o"}},
				},
				Action: initProject,
			},
			{
				Name: "make:migration",
				Aliases: []string{"m:m"},
				Usage: "fast make migration files",
				UsageText: "goof make:migration [table]",
				Action: makeMigration,
			},
			{
				Name: "db:migrate",
				Aliases: []string{"d:m"},
				Usage: "migration database files",
				UsageText: "goof db:migrate",
				Action: databaseMigrate,
			},
		},
	}

	// set options
	runError := app.Run(os.Args)
	if runError != nil {
		log.ErrorFatal(runError.Error())
	}
}



// init project, publish foundation files
func initProject(context *cli.Context) error {

	log.Info("publish files to:" + workDir)

	isOutput := context.Bool("output")
	names := AssetNames()

	for _, name := range names {

		info, err := AssetInfo(name)
		if err != nil && isOutput {
			log.ErrorFatal(err.Error())
		}

		contents, err := Asset(name)
		if err != nil && isOutput {

			log.ErrorFatal(err.Error())
		}

		dir := path.Dir(info.Name())
		filename := info.Name()

		if dir != "." {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil && isOutput {
				log.Error(err.Error())
			}
		}

		err = ioutil.WriteFile(filename, contents, info.Mode())
		if isOutput {
			if err != nil {
				log.ErrorFatal(err.Error())
			} else {
				log.Info("generate:" + info.Name())
			}
		}
	}

	log.Info("publish success")

	return nil
}

// make migration files
func makeMigration(context *cli.Context) error {

	// generate
	table := context.Args().First()

	now := time.Now().Format("20060102150405")
	packName := table + now
	filename := "database/migrations/" + now + table + ".go"

	err := ioutil.WriteFile(filename, []byte(stub(packName)), os.ModePerm)
	if err == nil {
		log.Info("publish migration to:" + filename)
	}

	return err
}

func databaseMigrate(context *cli.Context) error {

	// database/migrations path all files migrate
	rootPath := "database/migrations"

	log.Info("start migrate")

	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		if info == nil {
			return err
		}

		if info.IsDir() {
			return nil
		}


		return nil
	})
}



func stub(packName string) string {

	stub := `package %s

import (
	"github.com/jinzhu/gorm"
	"github.com/seth-shi/goof"
)

func init()  {

	goof.RegisterMigrate(Up, Down)
}


func Up(db gorm.DB)  {
	
}

func Down(db gorm.DB) {
	
}
`

	return fmt.Sprintf(stub, packName)
}