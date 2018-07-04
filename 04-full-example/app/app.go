package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ellisonleao/rest-api-go/04-full-example/db"
	"github.com/ellisonleao/rest-api-go/04-full-example/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

// App é a struct principal da nossa aplicação, que contém as dependencias necessarias para roda-la
type App struct {
	DB     db.Transport
	Router *mux.Router
}

// NewApp cria a nossa aplicacao com suas dependencias
func NewApp() (*App, error) {
	mongoURL := os.Getenv("MONGO_URL")
	mongo, err := db.NewMongoTransport(mongoURL)
	if err != nil {
		return nil, err
	}
	app := &App{
		DB: mongo,
	}

	app.MakeRoutes()
	return app, nil
}

// MakeRoutes cria as rotas da nossa api
func (a *App) MakeRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/players", a.listPlayers).Methods("GET")
	router.HandleFunc("/player", a.insertPlayer).Methods("POST")
	router.HandleFunc("/player/{id}", a.getPlayer).Methods("GET")
	router.HandleFunc("/player/{id}", a.deletePlayer).Methods("DELETE")

	a.Router = router
}

func (a *App) listPlayers(w http.ResponseWriter, r *http.Request) {
	var players []*models.Player
	// estamos dizendo que o retorno será em json
	w.Header().Set("Content-Type", "application/json")

	err := a.DB.FindAll(&players)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
	return
}

func (a *App) getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var player *models.Player

	// Tentamos pegar o player pelo ID no banco
	err := a.DB.FindByID(vars["id"], &player)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, "Player not Found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
	return
}

func (a *App) insertPlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	// validando o body request

	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid body request", http.StatusBadRequest)
		return
	}

	player.ID = bson.NewObjectId()

	// Poderiamos colocar mais validacoes aqui, mas essa nao era a intenção da talk..

	err = a.DB.Insert(player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// estamos dizendo que o retorno será em json
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&player)

}

func (a *App) deletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Tentamos pegar o player pelo ID no banco e remove-lo
	err := a.DB.Delete(vars["id"])
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, "Player not Found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
