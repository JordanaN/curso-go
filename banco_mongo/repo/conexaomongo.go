package repo

import (
	"gopkg.in/mgo.v2"
)

//Armazena a conexao com o Mongo
var SessaoMongo *mgo.Session

//AFaz a conexao com o Mongo
func AbreSessaoComMongo() (err error) {
	err = nil
	SessaoMongo, err = mgo.Dial("mongodb://localhost:27017/curso-go")
	return
}
