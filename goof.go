//go:generate go-bindata -prefix framework -o framework.go framework/...

package main

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"github.com/seth-shi/goof/utils"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
		Version: "v1.0.0",
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

	model, err := decodeMigrationName(table)
	if err != nil {
		return err
	}

	filename := "database/migrations/" + model.getFileName()
	var contents string
	switch model.Action {

	case MigrationCreateAction:
		contents = model.createMigrationTemplate()
	case MigrationUpdateAction:
		contents = model.updateMigrationTemplate()
	case MigrationDeleteAction:
		contents = model.deleteMigrationTemplate()
	}

	err = ioutil.WriteFile(filename, []byte(contents), os.ModePerm)
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

		model, err := decodeMigrationName(info.Name())
		if err != nil {
			return err
		}

		// 得到结构体的名字
		structName := model.getStructName()
		valType := reflect2.TypeByName("migrations." + structName)
		fmt.Println("migrations." + structName, valType.String())

		return nil
	})
}