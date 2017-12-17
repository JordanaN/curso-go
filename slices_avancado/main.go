package main

import (
	"fmt"
)

func main() {

	cidades := make([]string, 5)
	cidades[0] = "São Paulo"
	cidades[1] = "Curitiba"
	cidades[2] = "Porto Alegre"
	cidades[3] = "Rio de Janeiro"
	cidades[4] = "Palmas"
	fmt.Println(cidades, len(cidades), cap(cidades))

	//intervalos
	// Primeiro começa em 0 e segundo começa em 1
	capitaisBrasil := cidades[3:5]
	fmt.Println(capitaisBrasil)
	temp1 := cidades[:2]
	fmt.Println(temp1)
	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)
	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)

}
