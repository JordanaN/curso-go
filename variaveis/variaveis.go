package main

//variaveis sem uso podem impedir a compilação do código

import (
	"fmt"
)

//Tipos de variáveis basicas

var endereco string // ""
var telenone int //0
var comprou bool //false
var valor float64 // 0.00

//maneira compacta de declarar variaveis
var (
	nome string
	idade int
	palavra rune
)

//variaveis podem ser publicas ou privadas a diferença é a 1º letra da variavel
var estado string //privada
var Cidade string //publica

func main(){
	//outro modo de criar uma variavel sem utilizar o "var"
	local_trabalho := "Eduzz"
	fmt.Printf("Local de Trabalho: %s \n", local_trabalho)

}
