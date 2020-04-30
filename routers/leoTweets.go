package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Julianrt/twittor/bd"
)

//LeoTweets Endpoint para ver los tweet de un usuario
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id param needed", http.StatusBadRequest)
		return
	}

	paginaStr := r.URL.Query().Get("pagina")
	if len(paginaStr) < 1 {
		http.Error(w, "pagina param needed", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(paginaStr)
	if err != nil {
		http.Error(w, "pagina param have to be bigger than 0", http.StatusBadRequest)
		return
	}
	paginaI64 := int64(pagina)
	tweets, seObtuvieronTweets := bd.LeoTweets(ID, paginaI64)
	if !seObtuvieronTweets {
		http.Error(w, "error while trying to read tweets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
