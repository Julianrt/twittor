package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Julianrt/twittor/bd"
)

//ListaUsuarios muestra la lista de los usuarios
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pag, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Parámetro página debe ser entero mayor a 0", http.StatusBadRequest)
		return
	}

	results, status := bd.LeoUsuariosTodos(IDUsuario, int64(pag), search, typeUser)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
