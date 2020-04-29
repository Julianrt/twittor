package main

import (
	"log"

	"github.com/Julianrt/twittor/bd"
	"github.com/Julianrt/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}
