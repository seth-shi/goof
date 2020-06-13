package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// init project, publish foundation files
func initProject(context *cli.Context) error {

	dest, err := os.Getwd()
	if err != nil {
		return err
	}

	log.Println("publish files to:"+ dest)

	output := context.Bool("output")

	// 开始 copy 文件
	for _, file := range GetFrameworkFiles() {

		tmpName := dest + "/" + file.Name()

		if file.IsDir() {

			err := os.Mkdir(tmpName, os.ModePerm)

			if output {

				if err == nil {
					log.Println("publish success:" + tmpName)
				} else {
					log.Println(err.Error())
				}
			}
		}
	}

	return err
}