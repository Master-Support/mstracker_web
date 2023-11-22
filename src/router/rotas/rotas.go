package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}


func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasEnvio

	for _, rota := range rotas{
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return router
}