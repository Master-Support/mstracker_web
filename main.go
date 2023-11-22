package main

import (
	"fmt"
	"log"
	"mstracker_web/src/router"
	"mstracker_web/src/utils"
	"net/http"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}