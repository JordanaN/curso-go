package manipulador

import "html/template"

//faz o parser do arquivo html
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
