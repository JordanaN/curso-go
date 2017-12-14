package main

import "fmt"

//struct é uma coleção de dados
type Imovel struct {
	X int
	Y int
	Nome string
	valor int
}

func main(){
	//maneiras de se declarar e popular a structs

	casa := Imovel{}
	fmt.Printf("A casa é %+v\r\n", casa)

	//quando declado assim em uma linha os valores devem seguir a mesma ordem da declaração da struct
	apartamento := Imovel {45, 60, "Meu Apto", 100000}
	fmt.Printf("O apartamento é %+v\r\n", apartamento)

	//não precisa estar em ordem
	chacara := Imovel{
		X : 50,
		Y : 100,
		Nome : "Minha Chacara",
		valor : 60000,
	}

	fmt.Printf("A chacara é %+v\r\n", chacara)

	casa.Nome = "Minha casa"
	casa.valor = 50000
	casa.X = 30
	casa.Y = 50
	fmt.Printf("A casa é %+v\r\n", casa)	
}
