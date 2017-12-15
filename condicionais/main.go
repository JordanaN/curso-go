package main

import (
	"fmt"
)

func main() {
	meses := 6
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Esta devendo!")
	}

	if !situacao {
		fmt.Println("Não esta devendo!")
	}

	if cidade == "Teste" {
		fmt.Printf("O cliente vive na cidade: %s\r\n", cidade)
	}

	//condicional com atribuição de função e variavel, a variavel criada nesse if só vale para ele
	if descricao, status := tempoDeDivida(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}

	fmt.Println("Obrigada por nos consultar")

}

func tempoDeDivida(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo"
		return
	}

	descricao = "O cliente não esta devendo"
	return
}
