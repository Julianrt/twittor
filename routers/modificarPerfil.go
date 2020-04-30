package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/models"
)

//ModificarPerfil Endpoint para actualizar los datos del perfil de un usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Datos incorrectos -> "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := bd.ModificoUsuario(usuario, IDUsuario)
	if err != nil {
		http.Error(w, "error al identificar el registro -> "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se modifico el registro", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
