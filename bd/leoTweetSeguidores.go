package bd

import (
	"context"
	"time"

	"github.com/Julianrt/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
)

//LeoTweetSeguidores trae los tweets de las personas a las que se siguen
func LeoTweetSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	//-1 trae los tweet de forma descendentes por la fecha y 1 los trae de forma ascendente
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	var result []models.DevuelvoTweetsSeguidores
	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return result, false
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
