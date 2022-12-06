package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"net/internal/data"
	"os"
	"time"
)

const version = "1.0.0"

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
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config
	// read config from command
	flag.IntVar(&cfg.port, "port", 4000, "api server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:password@localhost:5432/greenlight?sslmode=disable", "PostgreSql DSN")

	flag.IntVar(&cfg.db.maxIdleConns, "db-max-open-conns", 25, "PostgresSql max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idel-conns", 25, "Postgresql max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "postgresSql max connection idle time")
	flag.Parse()
	// init logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Printf("database connection pool established")
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}
	// http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
