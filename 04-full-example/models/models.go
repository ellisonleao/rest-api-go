package models

import "github.com/globalsign/mgo/bson"

// Player é a struct que representa um jogador da nossa API
type Player struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Position string        `json:"position" bson:"position"`
}
