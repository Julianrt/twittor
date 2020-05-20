package models

//Relacion modelo para grabar la relacion de un usuario con otro
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuario_id"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuario_relacion_id"`
}
