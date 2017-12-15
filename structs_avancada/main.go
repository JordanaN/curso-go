package main

import "curso-go/structs_avancada/model"
import "fmt"
import "encoding/json"

func main() {

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
