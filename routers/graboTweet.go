package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/models"
)

//GraboTweet permite grabar el tweet en la BD
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "error al insertar el registro -> "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "no se logro insertar el registro ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
