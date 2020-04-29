package bd

import (
	"github.com/Julianrt/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin realiza el chequeo de login de la BD
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usuario, false
	}

	//CompareHashAndPassword pide un []byte de la password del usuario y
	//otro []byte de la password ingresada para compararlas y retorna un error si no coinciden
	err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		return usuario, false
	}
	return usuario, true
}
