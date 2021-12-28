package main

import (
	"errors"
	"fmt"
)

const (
	categoriaA = "A"
	categoriaB = "B"
	categoriaC = "C"
)

var minutosTrabajados int
var horasTrabajadas int

func calcularHoras(minutosTrabajados int) int {
	return minutosTrabajados / 60
}

func calcularSalario(categoria string, horasTrabajadas int) (float64, error) {
	var salarioFinal float64
	switch {
	case categoria == categoriaA:
		salarioFinal = float64(horasTrabajadas * 3000)
		salarioFinal += salarioFinal * 0.5
	case categoria == categoriaB:
		salarioFinal = float64(horasTrabajadas * 1500)
		salarioFinal += salarioFinal * 0.2
	case categoria == categoriaC:
		salarioFinal = float64(horasTrabajadas * 1000)
	default:
		return 0, errors.New("Categoría incorrecta")
	}
	return salarioFinal, nil
}

func main() {
	minutosTrabajados = 600
	horasTrabajadas = calcularHoras(minutosTrabajados)
	salarioFinal, err := calcularSalario(categoriaB, horasTrabajadas)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salario:", salarioFinal)
	}
}
