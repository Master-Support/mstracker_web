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
	rotas = append(rotas, rotasEnvio...)

	for _, rota := range rotas{
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	return router
}