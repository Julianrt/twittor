package bd

import (
	"context"
	"log"
	"time"

	"github.com/Julianrt/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion consulta la relacion entre dos usuarios
func ConsultoRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         relacion.UsuarioID,
		"usuariorelacionid": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}
