package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/models"
)

//ConsultaRelacion ve si hay relacion entre dos usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = id

	var respuesta models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(relacion)
	if err != nil || status == false {
		respuesta.Status = false
	} else {
		respuesta.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
