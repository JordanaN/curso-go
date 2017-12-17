package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"curso-go/servidor_web/model"
)

func Ola(w http.ResponseWriter, r *http.Request) {

	hora := time.Now().Format("00:00:00")
	pagina := model.Pagina{}
	pagina.Hora = hora
	//executando o template
	if err := Modelos.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
	}
}