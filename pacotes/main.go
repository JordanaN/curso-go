package main

import (
	"fmt"
	"../pacotes/prefixo"
	"../pacotes/operadora"
	//"pacotes/operadora"
	//"pacotes/prefixo"
)

var Username = "Jordana"

func main(){
	fmt.Printf("Nome do usuario: %s \r\n", Username)
	fmt.Printf("Prefixo da Cidade: %d\r\n", prefixo.Cidade)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Nome e Numero da Operadora: %s\r\n", operadora.PrefixoDeMariliaOperadora)
	fmt.Printf("Teste de variavel publica e privada: %s\r\n", prefixo.TestePrefixo)

}