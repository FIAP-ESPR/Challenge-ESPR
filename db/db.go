package db

import (
	"database/sql"
	"fiap/ancora/helper"
	"fmt"

	_ "github.com/lib/pq"
)

var postgres *sql.DB

func ConectPostgres() *sql.DB {
	conf := helper.GetConfig()

	conexao := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUsername, conf.DBPassword, conf.DBName)

	// fmt.Println(conexao)

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GetDB() *sql.DB {
	if postgres == nil {
		postgres = ConectPostgres()
	}
	return postgres
}

func InitDB() {
	postgres = ConectPostgres()
}
