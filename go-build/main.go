package main

import "curso-go/go-build/model"
import "fmt"
import "encoding/json"

func main() {
	//GO É COMPILADO!!!
	//comando go build cria um executavel dentro na pasta do meu projeto
	//renomear executavel go build -o novonome
	//GOOS - referencia o SO


	casa := model.Imovel{}
	casa.Nome = "Casa Rosa"
	casa.X = 45
	casa.Y = 20
	casa.SetValor(150000)

	fmt.Printf("Casa: %+v\r\n", casa)
	fmt.Printf("Valor da casa: %+v\r\n", casa.GetValor())

	//utilizando JSON
	objJSON, _ := json.Marshal(casa)

	fmt.Println("Casa em JSON: ", string(objJSON))

}