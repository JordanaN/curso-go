package main

import (
	"fmt"
	"net/http"

	"curso-go/servidor_web/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola Jordana!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)

	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Iniciou...")

	http.ListenAndServe(":8181", nil)
}
