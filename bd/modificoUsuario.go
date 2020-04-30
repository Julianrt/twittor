package bd

import (
	"context"
	"time"

	"github.com/Julianrt/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModificoUsuario permite modificar el perfil del usuario
func ModificoUsuario(usuario models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}
	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}
	registro["fechaNacimiento"] = usuario.FechaNacimiento
	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}
	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}
	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}
	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}
	if len(usuario.SitioWeb) > 0 {
		registro["sitioweb"] = usuario.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return false, err
	}

	filtro := bson.M{
		"_id": bson.M{
			"$eq": objID,
		},
	}

	_, err = col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
