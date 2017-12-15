package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3

	fmt.Print("O numero: ", numero, " se escreve assim ")

	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("tres.")
	}

	fmt.Print("Voce é da família do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough // segue para o proximo
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println(" Sim.")
	default:
		fmt.Println("Deixa pra lá")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("OK! fique até mais tarde")
	default:
		fmt.Println("Amanhã é dia de trabalho!")
	}

	numero = 150
	fmt.Println("Esse numero cabe num dígito?")
	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Tente dois digítos...")
	case numero >= 100:
		fmt.Println("Numero muito grande!")
	}

}
