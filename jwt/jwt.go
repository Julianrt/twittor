package jwt

import (
	"time"

	"github.com/Julianrt/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genera el encriptado con JWT
func GeneroJWT(usuario models.Usuario) (string, error) {

	miClavePrivada := []byte("clave_perrona_segura")

	payload := jwt.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Nombre,
		"apellidos":        usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia":        usuario.Biografia,
		"ubicacion":        usuario.Ubicacion,
		"sitio_web":        usuario.SitioWeb,
		"_id":              usuario.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClavePrivada)
	return tokenStr, err
}
