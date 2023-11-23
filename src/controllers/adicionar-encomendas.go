package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Conexão com a API
func AdicionarEncomenda(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	objeto, err := json.Marshal(map[string]string{
		"codigoObjeto": r.FormValue("codigoObjeto"),
		"nomeObjeto": r.FormValue("nomeObjeto"),
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:8080/objetos-correios/", "application/json", bytes.NewBuffer(objeto))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)
}
