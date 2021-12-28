package main

import "fmt"

var sueldo float64

func impuesto(sueldo float64) float64 {
	var resultado float64

	if sueldo > 50000 && sueldo <= 150000 {
		resultado = sueldo * 0.17
	} else if sueldo > 150000 {
		resultado = sueldo * 0.1
	}

	return resultado
}

func main() {
	sueldo = 50000
	fmt.Println("El impuesto para el sueldo de", sueldo, "es de:", impuesto(sueldo))
}
