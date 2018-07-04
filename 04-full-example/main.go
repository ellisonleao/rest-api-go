package main

import (
	"log"
	"net/http"

	"github.com/ellisonleao/rest-api-go/04-full-example/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatal("Could not start app")
	}
	log.Print("Starting server..")
	log.Fatal(http.ListenAndServe(":8000", app.Router))
}
