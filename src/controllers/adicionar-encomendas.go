package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//Conexão com a API
func AdicionarEncomenda(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	objeto, err := json.Marshal(map[string]string{
		"nomeObjeto": r.FormValue("nomeObjeto"),
		"codigoObjeto": r.FormValue("codigoObjeto"),
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

func VizualizarEncomenda(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    response, err := http.Get("http://localhost:8080/objetos-correios/objetos")
    if err != nil {
        http.Error(w, "Erro ao obter dados da API", http.StatusInternalServerError)
        log.Println("Erro ao obter dados da API:", err)
        return
    }
    defer response.Body.Close()

    var buffer bytes.Buffer

    _, err = io.Copy(&buffer, response.Body)
    if err != nil {
        http.Error(w, "Erro ao processar dados da API", http.StatusInternalServerError)
        log.Println("Erro ao processar dados da API:", err)
        return
    }

    // Defina o cabeçalho Content-Type como application/json
    w.Header().Set("Content-Type", "application/json")

    // Escreva os dados na resposta HTTP
    w.Write(buffer.Bytes())
}