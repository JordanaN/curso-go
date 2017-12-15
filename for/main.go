package main

import (
	"fmt"
)

func main() {

	//Tipos de for

	for i := 0; i < 10; i++ {
		fmt.Println("Qual o valor de i?", i)
	}

	valor := 0
	teste := true

	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("nº: ", valor)
	}

	for {
		valor--
		fmt.Println("nº: ", valor)
		if valor == 0 {
			break
		}
	}

	texto := "Aprendendo Golang"
	for i, letra := range texto {
		fmt.Printf("Texto [%d] = %q\r\n", i, letra)
	}

}
