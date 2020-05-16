package routers

import (
	"net/http"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/models"
)

//BajaRelacion realiza el borrado de la relacion entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(relacion)
	if err != nil {
		http.Error(w, "Error: borrar relacion -> "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "No se borro la relacion", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
