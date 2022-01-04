package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

	datosJson, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	productos := Productos{}
	err = json.Unmarshal(datosJson, &productos)
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/productos", func(ctxt *gin.Context) {
		ctxt.JSON(http.StatusOK, productos)
	})

	router.Run(":8080")
}
