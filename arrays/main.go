package main

import (
	"fmt"
)

func main() {
	var teste [3]int

	teste[0] = 2
	teste[1] = 1
	teste[2] = 0
	fmt.Println("Capaciadade do array: ", len(teste))

	cantores := [2]string{"Michael Jackson", "John Lennon"}
	fmt.Printf("Conteúdo do array: \r\n%v\r\n", cantores)

	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("Capacidade do array: ", len(capitais))
	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d] é %s\r\n", indice, cidade)
	}

}
