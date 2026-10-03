package models

import (
	"database/sql"
	"log"
	"os"
)

type DBModel struct {
	DB       *sql.DB
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewDBModel(db *sql.DB) *DBModel {
	il := log.New(os.Stdout, "\x1b[32mDB_INFO:\x1b[0m\t", log.Ldate|log.Ltime)
	errl := log.New(os.Stdout, "\x1b[31mDB_ERROR:\x1b[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &DBModel{
		DB:       db,
		infoLog:  il,
		errorLog: errl,
	}
}
