package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CopyDir(src string, dest string, force bool) error {

	err := filepath.Walk(src, func(file string, info os.FileInfo, err error) error {

		if info == nil {
			return err
		}

		fmt.Println(file)
		return nil
	})

	return err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}