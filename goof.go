package main

import (
	"github.com/seth-shi/goof/utils"
	"github.com/urfave/cli/v2"
	"os"
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
		Version: "v1.0.2",
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
		},
	}

	// set options
	runError := app.Run(os.Args)
	if runError != nil {
		log.ErrorFatal(runError.Error())
	}
}