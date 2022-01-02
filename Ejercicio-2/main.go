package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Producto struct {
	Nombre   string
	Id       int
	Precio   float64
	Cantidad int
}

func (p Producto) detalle() string {
	return fmt.Sprintf("%d\t\t\t\t\t\t%.2f\t%d", p.Id, p.Precio, p.Cantidad)
}

func crearProducto(nombre string, id int, precio float64, cantidad int) Producto {
	p := Producto{
		Nombre:   nombre,
		Id:       id,
		Precio:   precio,
		Cantidad: cantidad,
	}

	return p
}

func main() {
	data, err := os.Open("../myFileProducts.txt")
	// t := string(data)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ';' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma)

	for scanner.Scan() { // internally, it advances token based on sperator
		for _, v := range scanner.Text() {
			fmt.Println(string(v))
		}
		// fmt.Println(scanner.Text()) // token in unicode-char
		// fmt.Println(scanner.Bytes()) // token in bytes

	}

	// productos := []Producto{}

}
