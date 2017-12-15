package main

import (
	"curso-go/funcoes_basicas/matematica"
	"fmt"
)

func teste() {
	fmt.Printf("ola")
}

func main() {

	resultado := matematica.Calculo(matematica.Multiplicador, 2, 5)
	fmt.Printf("Resultado Multiplicação: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Soma, 2, 5)
	fmt.Printf("Resultado Soma: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisao, 2, 5)
	fmt.Printf("Resultado Divisao: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Subtracao, 2, 5)
	fmt.Printf("Resultado Subtracao: %d\r\n", resultado)

	resultado = matematica.Somaaaa(8, 5)
	fmt.Printf("Resultado da Somaaaa: %d\r\n", resultado)

	resultado, resto := matematica.DivisaoComResto(9, 4)
	fmt.Printf("Resultado da Divisão: %d\r\ne resultado com resto:  %d\r\n", resultado, resto)
}
