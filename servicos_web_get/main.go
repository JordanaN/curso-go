package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//sempre configurar o timeout
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")

	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do google: ", err.Error())
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)

		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina do google: ", err.Error())
			return
		}
		//imprimindo conteudo
		fmt.Println(string(corpo))
	}

	//usando o request com autenticação
	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro no request para o google: ", err.Error())
		return
	}
	//token
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do google: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina do google: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
