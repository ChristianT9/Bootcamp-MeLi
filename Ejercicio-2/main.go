package main

import "fmt"

func main() {
	var precio float64 = 126000
	var descuento float64 = 0.6

	precioFinal := precio * descuento

	fmt.Println(precioFinal)

}
