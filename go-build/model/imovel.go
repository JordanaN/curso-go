package model

/*
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}*/

//usando JSON
type Imovel struct {
	X     int `json:"coordenada_x"`
	Y     int `json:"coordenada_y"`
	Nome  string `json:"nome"`
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
