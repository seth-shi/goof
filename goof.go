//go:generate go-bindata -prefix framework -o framework.go framework/...

package main

import (
	"github.com/seth-shi/goof/utils"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path"
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
		if dir != "." {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil && isOutput {
				log.Error(err.Error())
			}
		}

		err = ioutil.WriteFile(name, contents, info.Mode())
		if isOutput {
			if err != nil {
				log.ErrorFatal(err.Error())
			} else {
				log.Info("generate:" + info.Name())
			}
		}
	}

	log.Info("publish success")

	return err
}