package main

import (
	"curso-go/Mapas/model"
	"fmt"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 8)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 25
	casa.Y = 30
	casa.SetValor(150000)

	//colocando no cache
	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim: %+v\r\n", imovel)

	}

	apto := model.Imovel{}
	apto.Nome = "Meu apato"
	apto.X = 10
	apto.Y = 45
	apto.SetValor(100000)

	//colocando no cache
	cache[apto.Nome] = apto //valor pode ser qq coisa

	fmt.Println("qtos itens tem no cache?\n", len(cache))

	//1º chave e o 2º objeto
	for chave, imovel := range cache {
		fmt.Printf("Chave: [%s] : %v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	// '_' representa que o dado não será usado o go ignora 
	_, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, "Casa Amarela")
	}

	fmt.Println("qtos itens tem no cache?\n", len(cache))

	
	
}
