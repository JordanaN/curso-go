package operadora

import (
	"curso-go/pacotes/prefixo"
	"strconv"
)

var NomeOperadora = "Tim"

var PrefixoDeMariliaOperadora = strconv.Itoa(prefixo.Cidade) + " " + NomeOperadora
