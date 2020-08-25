package main

import (
	"github.com/joho/godotenv"
)

type Env struct {

	workDir string
	migrationDir string
}

func (env *Env) loadConfig()  {

	var envs map[string]string

	envs, _ = godotenv.Read(env.workDir + "/.env")

	migrationDir, exists := envs["MIGRATION_DIR"]
	if ! exists {
		migrationDir = "database/migrations"
	}
	env.migrationDir = migrationDir
}