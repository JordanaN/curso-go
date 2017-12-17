package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[main] Erro ao abrir o arquivo: ", err.Error())
		return
	}

	/* varrer o arquivo
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println("O conteúdo é: ", linha)
	}*/

	//ler o arquivo
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Erro ao ler o arquivo: ", err.Error())
		return
	}
	// range para demonstrar dados do arquivo
	for indLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s\r\n", indLinha, linha)
		for indItem, item := range linha {
			fmt.Printf("Item[%d] é %s\r\n", indItem, item)
		}
	}

	arquivo.Close()
}
