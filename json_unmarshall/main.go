package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"curso-go/json_unmarshall/model"
)

func main() {
	//sempre configurar o timeout
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)

	if err != nil {
		fmt.Println("[main] Erro no request para servidor: ", err.Error())
		return
	}

	request.SetBasicAuth("teste", "teste")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do Servidor: ", err.Error())
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro no conteudo da pagina: ", err.Error())
			return
		}
		fmt.Println(" ")
		//Json Unmarshal
		post := model.BlogPost{}
		
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno JSON: ", err.Error())
			return
		}
		fmt.Printf("Post é: %+v\r\n", post)
	}
}
