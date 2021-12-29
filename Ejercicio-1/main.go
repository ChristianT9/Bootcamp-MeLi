package main

import "fmt"

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (e Estudiante) detalle() {
	fmt.Printf("Nombre:\t%s\nApellido:\t%s\nDNI:\t%s\nFecha:\t%s\n", e.Nombre, e.Apellido, e.DNI, e.Fecha)
}

func main() {
	e := Estudiante{"Christian", "Tamayo", "1113699092", "Diciembre 29 del 2021"}
	e.detalle()
}
