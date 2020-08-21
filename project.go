package main

import (
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path"
)

// init project, publish foundation files
func initProject(context *cli.Context) error {

	log.Info("publish files to:" + env.migrationDir)

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
