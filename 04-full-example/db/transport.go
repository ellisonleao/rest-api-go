package db

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Transport é uma interface de contrato para os metodos de banco de dados
type Transport interface {
	FindByID(ID string, data interface{}) error
	FindAll(data interface{}) error
	Insert(data interface{}) error
	Delete(ID string) error
}

// MongoTransport é a struct que vai implementar os metodos da interface Transport para o mongoDB
type MongoTransport struct {
	*mgo.Session
}

// NewMongoTransport cria um novo objeto MongoTransport usando uma mongo url
func NewMongoTransport(url string) (*MongoTransport, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoTransport{Session: session}, nil
}

// FindByID retorna uma busca de objeto por mongo ID
func (m *MongoTransport) FindByID(ID string, data interface{}) error {
	session := m.Session.Copy()
	defer session.Close()

	return session.DB("worldcup").C("players").FindId(bson.ObjectIdHex(ID)).One(data)
}

// FindAll retorna uma lista de objetos vindos do DB
func (m *MongoTransport) FindAll(data interface{}) error {
	session := m.Session.Copy()
	defer session.Close()

	return session.DB("worldcup").C("players").Find(bson.M{}).All(data)
}

// Insert adiciona um novo item no db
func (m *MongoTransport) Insert(data interface{}) error {
	session := m.Session.Copy()
	defer session.Close()

	return session.DB("worldcup").C("players").Insert(data)
}

// Delete remove um item no db
func (m *MongoTransport) Delete(ID string) error {
	session := m.Session.Copy()
	defer session.Close()

	return session.DB("worldcup").C("players").RemoveId(bson.ObjectIdHex(ID))
}
