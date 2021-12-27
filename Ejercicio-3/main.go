package main

import "fmt"

func main() {
	edad := 23
	empleado := true
	antiguedad := 2
	sueldo := 150000
	if edad > 22 && empleado && antiguedad > 1 {
		fmt.Println("El usuario cumplé con todos los requisitos\n")
		fmt.Println("Requisitos:\n")
		fmt.Println("- La edad del usuario es mayor a 22 años")
		fmt.Println("- El usuario se encuentra empleado")
		fmt.Println("- El usuario tiene más de un año de antigüedad en su trabajo\n")
		if sueldo > 100000 {
			fmt.Println("No se le cobrará interés por tener un sueldo mayor a $100.000\n")
		} else {
			fmt.Println("Se le cobrará interés por tener un sueldo menor a $100.000\n")
		}
	} else {
		fmt.Println("No cumplé con ninguno de los requisitos")
	}
}
