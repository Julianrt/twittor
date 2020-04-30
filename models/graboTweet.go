package models

import "time"

//GraboTweet estructura que tendra nuestro tweet en la BD
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"user_id,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
