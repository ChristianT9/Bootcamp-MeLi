package main

import "fmt"

func main() {
	var nombresEstudiantes = []string{"Benjamin", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Lista de nombres:\n")

	for _, item := range nombresEstudiantes {
		fmt.Println("°", item)
	}

	fmt.Println()

	nuevaLista := append(nombresEstudiantes, "Gabriela")

	fmt.Println("Se agrego a la nueva estudiante:", nuevaLista[len(nuevaLista)-1], "\n")

	for _, item := range nuevaLista {
		fmt.Println("°", item)
	}
}
