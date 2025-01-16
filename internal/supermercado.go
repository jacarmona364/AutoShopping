package internal

import (
	"errors"
	"fmt"
)

type Supermercado map[string][]Producto

func NuevoSupermercado(nombre string) (Supermercado, error) {
	if nombre == "" {
		return nil, fmt.Errorf("El supermercado debe tener un nombre")
	}
	return Supermercado{
		nombre: []Producto{},
	}, nil
}