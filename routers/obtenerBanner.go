package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Julianrt/twittor/bd"
)

//ObtenerBanner envia el banner al HTTP
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if len(userID) < 1 {
		http.Error(w, "Debe de enviar el parametro id", http.StatusBadRequest)
		return
	}

	userPerfil, err := bd.BuscoPerfil(userID)
	if err != nil {
		http.Error(w, "usuario no encontrado -> "+err.Error(), http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + userPerfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada (banner) -> "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "error al copiar la imagen (banner) -> "+err.Error(), http.StatusBadRequest)
		return
	}
}
