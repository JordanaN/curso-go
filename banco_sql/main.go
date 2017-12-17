package main

import (
	"fmt"
	"net/http"

	"curso-go/banco_sql/manipulador"
	"curso-go/banco_sql/repo"
)

func init() {
	//fmt.Println("Iniciando...")
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Jordana!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Iniciando...")
	http.ListenAndServe(":8181", nil)
}
