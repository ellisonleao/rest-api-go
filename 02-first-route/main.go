package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/players", getPlayers)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Players list here.")
}
