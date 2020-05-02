package routers

import (
	"net/http"

	"github.com/Julianrt/twittor/bd"
)

//EliminarTweet permite borrar un tweet determinado
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("id")
	if len(tweetID) < 1 {
		http.Error(w, "Debe de enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(tweetID, IDUsuario)
	if err != nil {
		http.Error(w, "error al borrar el tweet -> "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
