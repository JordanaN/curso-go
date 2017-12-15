package main

import "fmt"

type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	//maneira de declarar ponteiros new()
	casa := new(Imovel)
	fmt.Printf("A casa é: %+v\r\n", casa) // A casa é: &{X:0 Y:0 Nome: valor:0}

	//demonstra o endereço de memória da variavel/struct casa
	fmt.Printf("A casa é: %p\r\n %+v\r\n", &casa, casa)
	//A casa é: 0xc42000c028
	//&{X:0 Y:0 Nome: valor:0}

	sitio := Imovel{50, 60, "meu sitio", 150000}
	fmt.Printf("\nMeu sitio é: %p\r\n %+v\r\n", &sitio, sitio)

	mudarImovel(&sitio)

	fmt.Printf("\nMeu sitio é: %p\r\n %+v\r\n", &sitio, sitio)

}

// função para alterar dados direto pelo endereço da memoria a mesma nao usa return
func mudarImovel(imovel *Imovel) {
	imovel.X = imovel.X + 8
	imovel.Y = imovel.Y + 23

}
