package main

import "fmt"

const (
	pequeño = 1
	mediano = 2
	grande  = 3

	costoEnvioGrande = 2500.0
)

type tienda struct {
	Nombre    string
	Productos []producto
}

type producto struct {
	Tipo   int
	Nombre string
	Precio float64
}

type Producto interface {
	calcularCosto() float64
}

type Ecommerce interface {
	total() float64
	agregar(p producto)
}

func (p producto) calcularCosto() float64 {
	var costo float64
	switch p.Tipo {
	case pequeño:
		costo = p.Precio
	case mediano:
		costo = p.Precio + (p.Precio * 0.03)
	case grande:
		costo = p.Precio + (p.Precio * 0.06) + costoEnvioGrande
	}
	return costo
}

func (t tienda) total() float64 {
	var costoTotal float64
	for _, item := range t.Productos {
		costoTotal += item.calcularCosto()
	}
	return costoTotal
}

func (t *tienda) agregar(p producto) {
	t.Productos = append(t.Productos, p)
}

func nuevoProducto(tipoProducto int, nombre string, precio float64) producto {
	p := producto{
		Tipo:   tipoProducto,
		Nombre: nombre,
		Precio: precio,
	}
	return p
}

func nuevaTienda(nombre string) tienda {
	t := tienda{
		Nombre: nombre,
	}
	return t
}

func main() {
	tienda := nuevaTienda("Compañía S.A.")
	p1 := nuevoProducto(grande, "Ember WF", 270000)
	p2 := nuevoProducto(pequeño, "Goku", 100000)
	p3 := nuevoProducto(mediano, "Asta", 152000)
	p4 := nuevoProducto(grande, "Warframe videojuego", 200000)
	tienda.agregar(p1)
	tienda.agregar(p2)
	tienda.agregar(p3)
	tienda.agregar(p4)
	fmt.Printf("La tienda tiene los siguientes productos:\n\n")
	for _, item := range tienda.Productos {
		fmt.Printf("Nombre:\t%s\nTipo:\t%d\nPrecio:\t$%g\n\n", item.Nombre, item.Tipo, item.Precio)
	}
	fmt.Printf("El costo total de la tienda es de $%g\n\n", tienda.total())
}
