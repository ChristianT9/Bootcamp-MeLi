package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id            string  `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha"`
}

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(ctxt *gin.Context) {
		name := ctxt.Param("name")
		ctxt.String(http.StatusOK, "Hola %s", name)
	})

	router.Run(":8080")
}
