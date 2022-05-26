package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	conn *sql.DB
}

func GetConnection() (*DbConfig, error) {
	db, err := sql.Open("mysql", getStringConnection())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &DbConfig{
		conn: db,
	}, nil
}

func getStringConnection() string {
	envs := GetEnvs()
	return fmt.Sprintf("%s:%s@tcp(localhost:3307)/%s?charset=utf8&parseTime=true&loc=Local",
		envs.Db.User,
		envs.Db.Password,
		envs.Db.Name,
	)
}
