package internal

import (
	"errors"
	"fmt"
)

type Precio struct {
	Valor float64
}

func NuevoPrecio(valor float64) (Precio, error) {
	if valor < 0 {
		return Precio{}, fmt.Errorf("El precio no puede ser negativo")
	}
	return Precio{Valor: valor}, nil
}