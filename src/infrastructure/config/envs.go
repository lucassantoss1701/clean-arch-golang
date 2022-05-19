package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envs *Envs

type dbEnvs struct {
	User     string
	Password string
	Name     string
}

type Envs struct {
	Db *dbEnvs
}

func GetEnvs() *Envs {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if envs == nil {
		db := &dbEnvs{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		}

		return &Envs{
			Db: db,
		}
	}
	return envs
}
