package model

type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

//Sintaxe define que essa função pertence a struct Imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}

