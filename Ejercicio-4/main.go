package main

import "fmt"

func main() {
	/* 	Decidí hacerlo con un mapa, ya que hasta el momento no he tenido mucho contacto con esta estructura de dato */
	mes := 9
	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	fmt.Println(meses[mes])
}
