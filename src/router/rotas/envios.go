package rotas

import (
	"mstracker_web/src/controllers"
	"net/http"
)

var rotasEnvio = []Rota{
	{
		URI:    "/",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeObjetos,
	},
	{
		URI:    "/encomendas",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarTelaDeObjetos,
	},
	{
		URI:    "/adicionar-encomendas",
		Metodo: http.MethodPost,
		Funcao: controllers.AdicionarEncomenda,
	},
	{
		URI:    "/vizualizar-encomendas",
		Metodo: http.MethodGet,
		Funcao: controllers.VizualizarEncomenda,
	},
}
