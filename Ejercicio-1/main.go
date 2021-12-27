package main

import "fmt"

func main() {
	palabra := "Christian"
	fmt.Println("Cantidad de letras:", len(palabra))
	for _, v := range palabra {
		fmt.Println(string(v))
	}
}
