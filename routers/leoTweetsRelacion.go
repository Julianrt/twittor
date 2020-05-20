package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Julianrt/twittor/bd"
)

//LeoTweetsSeguidores handler: mostrará los tweets de los usuarios que sigo
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "el parámetro página debe de ser numérico entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "No se pudieron leer los tweets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
