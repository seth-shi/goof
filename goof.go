package main

import (
	"github.com/seth-shi/goof/utils"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	log = utils.GetLogInstance()
)

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

	dest, err := os.Getwd()
	if err != nil {
		return err
	}


	log.Info("publish files to:" + dest)


	//isOutput := context.Bool("output")

	return err
}