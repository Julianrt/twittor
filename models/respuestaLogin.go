package models

//RespuestaLogin tiene el token que se devuelve von el login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
