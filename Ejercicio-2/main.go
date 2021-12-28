package main

import (
	"errors"
	"fmt"
)

func promedio(calificaciones ...int) (int, error) {
	var suma int
	var promedio int

	for _, item := range calificaciones {
		if item < 0 {
			return 0, errors.New("No se admiten números negativos para calcular el promedio")
		}
		suma += item
	}

	promedio = suma / len(calificaciones)

	return promedio, nil
}

func main() {
	promedio, err := promedio(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	if err == nil {
		fmt.Println("El promedio de la calificaciones de los estudiantes es:", promedio)
	} else {
		fmt.Println(err)
	}
}
