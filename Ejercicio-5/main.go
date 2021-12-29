package main

import (
	"errors"
	"fmt"
)

var animales = map[string]float64{"Tarántula": 150, "Hamster": 250, "Perro": 10, "Gato": 5}

var nombreAnimal string

var cantidadAnimal int

func animal(animal string) (func(cantidad float64) float64, error) {
	switch animal {
	case "Perro":
		return perro, nil
	case "Hamster":
		return hamster, nil
	case "Gato":
		return gato, nil
	case "Tarántula":
		return tarantula, nil
	default:
		return nil, errors.New("Este tipo de animal no existe")
	}
}

func perro(cantidad float64) float64 {
	return animales["Perro"] * cantidad
}

func hamster(cantidad float64) float64 {
	return animales["Hamster"] * cantidad
}

func gato(cantidad float64) float64 {
	return animales["Gato"] * cantidad
}

func tarantula(cantidad float64) float64 {
	return animales["Tarántula"] * cantidad
}

func main() {
	nombreAnimal = "Perro"
	cantidadAnimal = 20
	function, err := animal(nombreAnimal)
	if err != nil {
		fmt.Println(err)
	} else {
		resultado := function(float64(cantidadAnimal))
		if nombreAnimal == "Perro" || nombreAnimal == "Gato" {
			resultado = resultado * 1000 // Hago esto para que todos los resultados queden en gramos, ya que la comida para Gatos y Perros esta dada en Kilogramos.
		}
		fmt.Println("La cantidad de comida para", cantidadAnimal, nombreAnimal+"s", "es de:", resultado, "Gramos.")
	}
}
