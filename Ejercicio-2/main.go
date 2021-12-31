package main

import "fmt"

type Matrix struct {
	matriz      [3][3]float64
	alto        float64
	ancho       float64
	cuadratica  bool
	valorMaximo float64
}

func (m *Matrix) set(valorMaximo float64, alto, ancho float64, matriz [3][3]float64) {
	m.matriz = matriz
	m.alto = alto
	m.ancho = ancho
	m.valorMaximo = valorMaximo
	if ancho == alto {
		m.cuadratica = true
	} else {
		m.cuadratica = false
	}
}

func (m Matrix) print() {
	fmt.Println("Es cuadrática:", m.cuadratica)
	for i := 0; i < len(m.matriz); i++ {
		for j := 0; j < len(m.matriz[i]); j++ {
			fmt.Printf("%v\t", m.matriz[i][j])
		}
		fmt.Println()
	}
}

func main() {
	const alto uint = 3
	const ancho uint = 3
	matrix := Matrix{}
	matriz := [alto][ancho]float64{{1, 2, 5}, {3, 4, 6}, {1, 1, 1}}
	matrix.set(10, float64(alto), float64(ancho), matriz)
	matrix.print()
}
