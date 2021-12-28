package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func operacion(operacion string) (func(valores ...float64) (float64, error), error) {
	switch operacion {
	case minimo:
		return calcularMinimo, nil
	case promedio:
		return calcularPromedio, nil
	case maximo:
		return calcularMaximo, nil
	default:
		return nil, errors.New("Operación inválida")
	}
}

func calcularMinimo(valores ...float64) (float64, error) {
	var resultado float64
	for i := 0; i < len(valores); i++ {
		if i == 0 {
			resultado = valores[i]
		} else if valores[i] < resultado {
			resultado = valores[i]
		}
	}
	return resultado, nil
}

func calcularPromedio(valores ...float64) (float64, error) {
	var suma float64
	for _, item := range valores {
		if item < 0 {
			return 0, errors.New("No se puede calcular el promedio con un número negativo")
		}
		suma += item
	}
	return suma / float64(len(valores)), nil
}

func calcularMaximo(valores ...float64) (float64, error) {
	var resultado float64
	for i := 0; i < len(valores); i++ {
		if i == 0 {
			resultado = valores[i]
		} else if valores[i] > resultado {
			resultado = valores[i]
		}
	}
	return resultado, nil
}

func main() {
	minFunc, errMin := operacion(minimo)
	promFunc, errMax := operacion(promedio)
	maxFun, errProm := operacion(maximo)

	if errMin == nil {
		valorMinimo, err := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		if err == nil {
			fmt.Println("El número mínimo es:", valorMinimo)
		} else {
			fmt.Println("A ocurrido un error al calcular el mínimo.", err)
		}
	} else {
		fmt.Println(errMin)
	}

	if errMax == nil {
		valorMaximo, err := maxFun(2, 3, 3, 4, 1, 2, 4, 5)
		if err == nil {
			fmt.Println("El número máximo es:", valorMaximo)
		} else {
			fmt.Println("A ocurrido un error al calcular el máximo.", err)
		}
	} else {
		fmt.Println(errMax)
	}

	if errProm == nil {
		valorPromedio, err := promFunc(2, 3, 3, 4, 1, -2, 4, 5)
		if err == nil {
			fmt.Println("El promedio es:", valorPromedio)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(errProm)
	}

}
