package main

import (
	"fmt"
	"os"
)

type Producto struct {
	Nombre   string
	Id       int
	Precio   float64
	Cantidad int
}

func (p Producto) detalle() string {
	return fmt.Sprintf("Nombre:\t%s; Id:\t%d Precio:\t$%.2f; Cantidad:\t%d;\n", p.Nombre, p.Id, p.Precio, p.Cantidad)
}

func main() {
	p := Producto{
		Nombre:   "Café con leche",
		Id:       00001,
		Precio:   8000,
		Cantidad: 2,
	}

	p2 := Producto{
		Nombre:   "Café negro",
		Id:       23901,
		Precio:   6500,
		Cantidad: 1,
	}

	p3 := Producto{
		Nombre:   "Capuccino",
		Id:       34032,
		Precio:   7500,
		Cantidad: 3,
	}

	d1 := []byte(p.detalle() + p2.detalle() + p3.detalle())

	err := os.WriteFile("./myFileProducts.txt", d1, 0644)
	if err != nil {
		fmt.Println("Hubo un error al crear el archivo")
	} else {
		fmt.Println("No hubo ningún error: ", err)
	}
}
