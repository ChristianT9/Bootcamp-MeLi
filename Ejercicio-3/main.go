package main

import (
	"fmt"
	"time"
)

func main() {
	err := fmt.Errorf("momento del error: %v", time.Now())
	fmt.Println("error ocurrido: ", err)
}
