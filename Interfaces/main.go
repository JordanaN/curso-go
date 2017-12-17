package main

import (
	"fmt"

	"curso-go/Interfaces/model"
)

func main() {

	ave := model.Ave{}
	ave.Nome = "Maria Silva"

	queroAcordarComUmCacarejo(ave)
	queroOuvirUmaPataNoLago(ave)
}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
