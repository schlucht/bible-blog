package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/bibleblog/cmd/application"
	"github.com/schlucht/bibleblog/cmd/routes"
	"github.com/schlucht/bibleblog/pkg/drivers"
)

func main() {
	port := 8088
	dsn := "./db/db.db?_journal_mode=WAL&_busy_timeout=5000&_synchronous=NORMAL&_cache_size=-64000&_foreign_keys=ON"

	session := scs.New()
	session.Lifetime = 24 * time.Hour

	db, err := drivers.SqlLite(dsn)
	if err != nil {
		fmt.Println(err)
	}
	app := application.NewApp(port, session, db)
	
	routes := routes.Routes(&app)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           routes,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.InfoLog.Printf("Starting server on port %s", srv.Addr)

	err = srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

}
