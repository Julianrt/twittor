package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Julianrt/twittor/bd"
)

//VerPerfil permite extraer los valores del perfil de un usuario
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe envier el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "No se encontro el usuario -> "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
