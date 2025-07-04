package main

import (
	"github.com/gonzaloan/go-awesome/internal/db"
	"github.com/gonzaloan/go-awesome/internal/env"
	"github.com/gonzaloan/go-awesome/internal/store"
	"log"
)

func main() {

	cfg := config{
		address: env.GetString("ADDR", ":8080"),
		dbConfig: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenCons: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleCons: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	database, err := db.New(cfg.dbConfig.addr, cfg.dbConfig.maxOpenCons, cfg.dbConfig.maxIdleCons, cfg.dbConfig.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	storage := store.NewStorage(database)

	app := &application{
		config: cfg,
		store:  storage,
	}
	log.Printf("Starting server on address %s", app.config.address)

	log.Fatal(app.listen(app.mount()))
}
