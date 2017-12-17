package repo

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"	
)

// armazena a conexão com o banco de dados
var Db *sqlx.DB

//Abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/curso-go")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
