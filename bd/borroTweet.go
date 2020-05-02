package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BorroTweet borra un tweet determinado
func BorroTweet(tweetID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	objID, err := primitive.ObjectIDFromHex(tweetID)
	if err != nil {
		return err
	}

	condicion := bson.M{
		"_id":    objID,
		"userid": userID,
	}

	_, err = col.DeleteOne(ctx, condicion)
	return err
}
