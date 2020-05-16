package routers

import (
	"net/http"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/models"
)

//AltaRelacion handler para empezar a seguir una cuenta
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro is es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Error al insertar relacion -> "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "Relacion no se registro ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
