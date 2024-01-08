package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bennett-matt/blog/internal/data"
	"github.com/bennett-matt/blog/internal/jsonlog"
	"github.com/jmoiron/sqlx"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	tc     *TemplateCache
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("BLOG_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max idle time")
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	templateCache, err := NewTemplateRender()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	// TODO: implement `openDB`
	// db, err := openDB(cfg)
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// defer db.Close()
	// logger.Println("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		// models: data.NewModels(db),
		tc: templateCache,
	}

	svr := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.PrintFatal(svr.ListenAndServe(), nil)
}

func openDB(config config) (*sqlx.DB, error) {
	return nil, nil
}

func (a *application) render(w http.ResponseWriter, r *http.Request, t Template) {
	template, ok := a.tc.templates[t.View]
	if !ok {
		a.logger.PrintInfo("not ok", nil)
	}

	if err := template.ExecuteTemplate(w, t.Layout, t.Data); err != nil {
		a.logger.PrintError(err, nil)
	}
}
