package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Player é a struct que representa um jogador da nossa API
type Player struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Height   float32 `json:"height"`
}

// Vamos supor que a lista de players está vindo de um DB
var players = []*Player{
	&Player{
		ID:       123,
		Name:     "Coutinho",
		Position: "ME",
		Height:   1.78,
	},
	&Player{
		ID:       456,
		Name:     "Neymar",
		Position: "PE",
		Height:   1.75,
	},
	&Player{
		ID:       789,
		Name:     "William",
		Position: "PD",
		Height:   1.75,
	},
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/players", listPlayers)
	router.HandleFunc("/players/{id}", getPlayer)

	log.Print("Starting server..")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func listPlayers(w http.ResponseWriter, r *http.Request) {

	// estamos dizendo que o retorno será em json
	w.Header().Set("Content-Type", "application/json")

	// transformamos o slice de players em JSON e mandamos o retorno
	json.NewEncoder(w).Encode(players)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, ok := vars["id"]; !ok {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	// convertendo o id para int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Id must be an int", http.StatusBadRequest)
		return
	}

	// vamos supor que estamos fazendo uma query no banco e buscando o player
	encoder := json.NewEncoder(w)
	for _, player := range players {
		if player.ID == id {
			w.Header().Set("Content-Type", "application/json")
			encoder.Encode(player)
			return
		}
	}

	http.Error(w, "Player not Found", http.StatusNotFound)
}
