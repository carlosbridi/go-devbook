package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacao = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
