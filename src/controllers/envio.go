package controllers

import (
	"mstracker_web/src/utils"
	"net/http"
)

func CarregarTelaDeObjetos(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "envios.html", nil)
}
