package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// init project, publish foundation files
func initProject(context *cli.Context) error {

	fmt.Println("publish files to:" + env.migrationDir)

	isOutput := context.Bool("output")

	for _, filename := range AssetNames() {

		contents, err := Asset(filename)
		if err != nil && isOutput {

			fmt.Println(err.Error())
			continue
		}

		dir := filepath.Dir(filename)
		if dir != "." {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil && isOutput {
				fmt.Println(err.Error())
				continue
			}
		}

		err = ioutil.WriteFile(filename, contents, 0755)
		if isOutput {
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("generate:" + filename)
			}
		}
	}

	fmt.Println("publish success")

	return nil
}
