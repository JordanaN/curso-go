package repo

import (
	"curso-go/banco_mongo/model"
	"gopkg.in/mgo.v2/bson"
)

//Retorna um local do MongoDB
func ObtemLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("local")
	err = colecao.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}

//Registra a consulta ao Local
func SalvaLog(reg model.RegistroLog) (err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("logvisitas")
	err = colecao.Insert(reg)
	return
}
