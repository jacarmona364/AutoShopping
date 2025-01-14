package internal

import "errors"

type Supermercado struct {
	Nombre string
}

func NuevoSupermercado(nombre string) (Supermercado, error) {
	if nombre == "" {
		return Supermercado{}, fmt.Errorf("El supermercado debe tener nombre")
	}
	return Supermercado{Nombre: nombre}, nil
}

type Inventario struct {
	Productos map[string][]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: make(map[string][]Producto),
	}
}