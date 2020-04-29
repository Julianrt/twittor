package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/jwt"
	"github.com/Julianrt/twittor/models"
)

//Login endpoint para manejar el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var usuario models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "No se pudo procesar el json -> "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(usuario.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", http.StatusBadRequest)
		return
	}
	documento, existe := bd.IntentoLogin(usuario.Email, usuario.Password)
	if !existe {
		http.Error(w, "Usuario y/o contraseña inválidos", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Error al generar el token -> "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//EXTRA
	//GENERAR una cookie en el navegador

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
