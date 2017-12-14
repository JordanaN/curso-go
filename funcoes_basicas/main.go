package main

import (
	"fmt"
	"../funcoes_basicas/matematica"
)

func main(){
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 5)
	fmt.Printf("Resultado Multiplicação: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Soma, 2, 5)
	fmt.Printf("Resultado Multiplicação: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisao, 2, 5)
	fmt.Printf("Resultado Multiplicação: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Subtracao, 2, 5)
	fmt.Printf("Resultado Multiplicação: %d\r\n", resultado)
}

