package model

import (
	"errors"
)

//usando JSON
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//Erro personalizado
var ErrValorImovel = errors.New("O valor do imóvel é invalido")
var ErrValorImovelMuitoAlto = errors.New("O valor do imóvel é muito alto")


//Sintaxe define que essa função pertence a struct Imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorImovel
		return
	}else if valor > 10000000{
		err = ErrValorImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
