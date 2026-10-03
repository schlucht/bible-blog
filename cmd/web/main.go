package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/bibleblog/cmd/web/config"
	"github.com/schlucht/bibleblog/cmd/web/frontend"
	"github.com/schlucht/bibleblog/pkg/drivers"
)

const (
	port = 5100
	env  = "development"
	// dsn  = "ots:goweb@/goweb?parseTime=true"
	// dsn = "ots:goweb@tcp(127.0.0.1:3306)/goweb?parseTime=true"
	dsn = "schmidschluch:regal@tcp(db51.hostpark.net)/schmidschluch?parseTime=true"
	// dsn = "schmidschluch4:Schlucht6@tcp(db8.hostpark.net)/schmidschluch4?parseTime=true"
)

func main() {
	app := config.AppConfig{
		UseCache: false,
	}

	tc := make(map[string]*template.Template)
	app.TemplateCache = tc
	app.MySqlLog = true

	app.InfoLog = log.New(os.Stdout, "\x1b[32mINFO:\x1b[0m\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "\x1b[31mERROR:\x1b[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	app.Session = session

	db, err := drivers.MySqlDB(dsn)
	if err != nil {
		app.ErrorLog.Printf("MySql Fehler: %s", err)
		app.MySqlLog = false
	}
	defer db.Close()

	frontend.NewRenderer(&app)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           routes(&app),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.InfoLog.Printf("Starting server on port %s", srv.Addr)

	// ---------------------------------------------------------
	// Server starten
	// ---------------------------------------------------------
	err = srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}
