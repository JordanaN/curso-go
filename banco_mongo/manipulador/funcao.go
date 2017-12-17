package manipulador

import (
	"fmt"
	"net/http"
)

func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "manipulador de função do pacote")
}
