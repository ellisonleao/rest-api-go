package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ellisonleao/rest-api-go/04-full-example/db"
	"github.com/ellisonleao/rest-api-go/04-full-example/models"
)

var (
	app   App
	mongo db.MockTransport
)

// func TestMain(m *testing.M) {
// }

func TestListPlayers(t *testing.T) {
	mongo := &db.MockTransport{}
	app := &App{
		DB: mongo,
	}
	app.MakeRoutes()
	var players []*models.Player

	mongo.On("FindAll", &players).Return(nil)

	req, err := http.NewRequest("GET", "http://localhost:8000/players/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	app.listPlayers(res, req)

	var result []map[string]interface{}
	playersList := json.NewDecoder(res.Body).Decode(&result)
	if playersList != nil {
		t.Fatal("Empty players on DB should return nil as response")
	}
}
