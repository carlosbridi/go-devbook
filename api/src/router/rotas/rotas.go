package rotas

import "net/http"

//Rota representa estrutura de Rota
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}
