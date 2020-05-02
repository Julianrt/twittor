package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/models"
)

//SubirAvatar endpoint para subir avatar al servidor
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	split := strings.Split(handler.Filename, ".")
	extensionArchivo := split[len(split)-1]
	rutaArchivoNuevo := "uploads/avatars/" + IDUsuario + "." + extensionArchivo

	f, err := os.OpenFile(rutaArchivoNuevo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error al crear el nuevo archivo del avatar -> "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error al copiar el archivo enviado -> "+err.Error(), http.StatusInternalServerError)
		return
	}

	var usuario models.Usuario
	usuario.Avatar = IDUsuario + "." + extensionArchivo
	status, err := bd.ModificoUsuario(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "error al guardar la ruta del avatar en la BD -> "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
