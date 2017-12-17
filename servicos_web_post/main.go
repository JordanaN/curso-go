package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"curso-go/servicos_web_post/model"
)

func main() {
	//sempre configurar o timeout
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Jordana Nogueira"

	//transformando em JSON
	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o JSON: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://requestb.in/1agagio1", bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[main] Erro ao enviar um request: ", err.Error())
		return
	}

	//autenticação
	request.SetBasicAuth("fizz", "buzz")

	//tipo de request
	request.Header.Set("content-type", "application/json; charset=utf-8")
	//Do executa a chamada
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o servico do request: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
