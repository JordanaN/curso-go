package main

import (
	"curso-go/tratamento_erro/model"
	"encoding/json"
	"fmt"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Rosa"
	casa.X = 45
	casa.Y = 20
	//if de uma linha com a atribuição de função
	if err := casa.SetValor(150000000); err != nil {
		fmt.Println("[main] Houve um erro na: ", err)
		if err == model.ErrValorImovelMuitoAlto {
			fmt.Println("Corretor, por favor verifique a sua avaliação!")
		}
		return
	}

	fmt.Printf("Casa: %+v\r\n", casa)
	fmt.Printf("Valor da casa: %+v\r\n", casa.GetValor())

	objJSON, err := json.Marshal(casa)
	//ponteiro de um objeto que implementa a interface de Error,
	//nil = nulo
	if err != nil {
		fmt.Println("[main] Houve um erro na geracao do objeto JSON: ", err.Error)
	}

	fmt.Println("Casa em JSON: ", string(objJSON))

}
