package manipulador

import "html/template"

//faz o parser do arquivo ola.html
var Modelos = template.Must(template.ParseFiles("html/ola.html"))

//faz o parser do arquivo local.html
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))