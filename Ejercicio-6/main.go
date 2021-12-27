package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Println("Edad de Benjamin:", employees["Benjamin"], "años")
	mayoresA21 := 0
	for _, item := range employees {
		if item > 21 {
			mayoresA21++
		}
	}
	fmt.Println("Empleados mayores a 21 años:", mayoresA21)
	employees["Federico"] = 25
	fmt.Println("Se agrego el empleado Federico el cual tiene", employees["Federico"], "años")
	delete(employees, "Pedro")
	fmt.Println("Se elimino a Pedro de la lista de empleados")
	fmt.Println("La lista final de empleados es la siguiente:")
	for key, item := range employees {
		fmt.Println("°", key, item, "años")
	}
}
