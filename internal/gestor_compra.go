package internal

import "errors"

type GestorCompra map[string][]Producto

func NuevoGestorCompra(nombre string) (GestorCompra, error) {
	return GestorCompra{
		nombre: []Producto{},
	}, nil
}