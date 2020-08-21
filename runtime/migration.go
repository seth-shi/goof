package goof

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/seth-shi/goof/utils"
	"os"
	"path/filepath"
	"runtime"
)

var fileMigrations []string
var registerMigrations = make(map[string]struct {
	Up, Down func(db *gorm.DB)
})

type migration struct {
	ID        uint `gorm:"primary_key"`
	Migration string
	Batch     uint
}

func RegisterMigrate(up, down func(db *gorm.DB)) {

	_, filename, _, _ := runtime.Caller(1)

	registerMigrations[filepath.Base(filename)] = struct{ Up, Down func(db *gorm.DB) }{Up: up, Down: down}
}

func LoadMigrationFiles(dirPath string) ([]string, error) {

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s directory does not exists", dirPath)
	}

	files, err := filepath.Glob(dirPath + "/**.go")
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		base := filepath.Base(file)
		if _, exists := registerMigrations[base]; ! exists {
			return []string{}, ErrNotRegisterMigration
		}

		fileMigrations = append(fileMigrations, base)
	}

	return fileMigrations, nil
}

func migratedFiles() (files []string, err error) {

	if err = orm.AutoMigrate(&migration{}).Error; err != nil {
		return
	}

	err = orm.Model(&migration{}).Pluck("Migration", &files).Error
	if err != nil {
		return
	}

	return
}

func migrationMaxBatch() (uint, error) {

	var max uint
	row := orm.Model(&migration{}).Select("MAX(Batch)").Row()
	if err := row.Scan(&max); err != nil {
		return max, err
	}

	return max, nil
}

func migrate() error {

	dbMigrations, err := migratedFiles()
	if err != nil {
		return err
	}

	diffFiles := utils.Intersect(fileMigrations, dbMigrations)
	if len(diffFiles) == 0 {
		fmt.Println("Nothing to migrate.")
		return nil
	}


	maxBatch, err := migrationMaxBatch()
	if err != nil {
		return err
	}

	orm.Begin()
	defer func() {

		if e := recover(); e != nil {
			orm.Rollback()
		} else {
			orm.Commit()
		}
	}()

	for _, file := range diffFiles {

		m, _ := registerMigrations[file]
		m.Up(orm)
		orm.Create(&migration{
			Migration: file,
			Batch: maxBatch,
		})
	}

	return err
}