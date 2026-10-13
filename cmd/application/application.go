package application

import (
	"log"
	"os"
	"database/sql"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Port     int
	Session  *scs.SessionManager
	DB *sql.DB
}

func NewApp(port int, session *scs.SessionManager, db *sql.DB) Application {
	app := Application{
		InfoLog:  log.New(os.Stdout, "\x1b[32mINFO:\x1b[0m\t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stdout, "\x1b[31mERROR:\x1b[0m\t", log.Ldate|log.Ltime|log.Lshortfile),
		Port:     port,
		Session:  session,
		DB: db,
	}
	return app
}
