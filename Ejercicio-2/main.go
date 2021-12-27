package main

import "fmt"

var humedad float64
var temperatura float64
var presion float64

func main() {
	temperatura = 23
	humedad = 81
	presion = 1014

	fmt.Println("Temperatura:", temperatura, "Cº")
	fmt.Println("Humedad:", humedad, "%")
	fmt.Println("Presión:", presion, "hPa")
}
