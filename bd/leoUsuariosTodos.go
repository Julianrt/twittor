package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Julianrt/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoUsuariosTodos lee los usuarios registrados en el sistema
func LeoUsuariosTodos(ID string, page int64, search, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var usuario models.Usuario
		err := cur.Decode(&usuario)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relacion models.Relacion
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = usuario.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(relacion)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if relacion.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			usuario.Password = ""
			usuario.Biografia = ""
			usuario.SitioWeb = ""
			usuario.Ubicacion = ""
			usuario.Banner = ""
			usuario.Email = ""

			results = append(results, &usuario)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
