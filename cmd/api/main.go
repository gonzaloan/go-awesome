package main

import (
	"github.com/gonzaloan/go-awesome/internal/env"
	"log"
)

func main() {

	app := &application{
		config: config{
			address: env.GetString("ADDR", ":8080"),
		},
	}
	log.Printf("Starting server on address %s", app.config.address)

	log.Fatal(app.listen(app.mount()))
}
