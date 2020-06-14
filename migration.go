package main

import (
	"errors"
	"github.com/iancoleman/strcase"
	"regexp"
	"strings"
	"time"
)

type MigrationAction int8

const (
	MigrationCreateAction MigrationAction = iota
	MigrationUpdateAction
	MigrationDeleteAction
)

type MigrationManager struct {

	existsMigrations map[string]string
}

type MigrationModel struct {

	Table string
	Action MigrationAction
	Field string
	CreatedAt time.Time
}

// name => YYYY_mm_dd_HH_mm_ss_action_table_
func decodeMigrationName(name string) (MigrationModel, error) {

	if isOk, _ := regexp.MatchString(`\d{4}_\d{2}_\d{2}_\d{2}_\d{2}_\d{2}`, name); isOk {

		name = name[20:]
	}

	var model MigrationModel
	model.Table = name
	model.CreatedAt = time.Now()

	chunkStr := strings.Split(name, "_")
	if len(chunkStr) == 0 {
		return model, errors.New("decode fail migration " + name)
	}


	switch chunkStr[0] {
	case "create":
		// create_users_table
		if len(chunkStr) < 3 {
			return model, errors.New("create migration name is create_tb_table")
		}
		model.Action = MigrationCreateAction
	case "update", "add":
		// add_field_to_table
		if len(chunkStr) < 5 {
			return model, errors.New("update migration name is add_field_to_tb_table")
		}
		model.Action = MigrationUpdateAction
		model.Field = chunkStr[1]
	case "del", "delete":
		// del_field_to_table
		if len(chunkStr) < 5 {
			return model, errors.New("delete migration name is del_field_to_tb_table")
		}
		model.Action = MigrationDeleteAction
		model.Field = chunkStr[1]
	default:
		return model, errors.New("migration file name start with create | add | del")
	}

	return model, nil
}


func (m MigrationModel) getStructName () string  {

	return strcase.ToCamel(m.Table)  + "_" + m.CreatedAt.Format("20060102150405")
}

func (m MigrationModel) getFileName () string  {

	return m.CreatedAt.Format("2006_01_02_15_04_05") + "_" + strcase.ToSnake(m.Table) + ".go"
}

func (m MigrationModel) createMigrationTemplate() string {

	tableName := m.getStructName()

	return `
package migrations

type `+ tableName +` struct {

}

func (`+ tableName +`) up()  {

}

func (`+ tableName +`) down()  {

}
`
}

func (m MigrationModel) updateMigrationTemplate() string {

	tableName := m.getStructName()

	return `
package migrations

type `+ tableName +` struct {

}

func (`+ tableName +`) up()  {

}

func (`+ tableName +`) down()  {

}
`
}

func (m MigrationModel) deleteMigrationTemplate() string {

	tableName := m.getStructName()

	return `
package migrations

type `+ tableName +` struct {

}

func (`+ tableName +`) up()  {

}

func (`+ tableName +`) down()  {

}
`
}