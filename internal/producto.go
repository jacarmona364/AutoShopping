package internal

import (
	"errors"
	"fmt"
)

type Producto struct {
	Nombre   string
	Precio   Precio
	Alergias []Alergia
}

func NuevoProducto(nombre string, precio Precio, alergias []Alergia) (Producto, error) {
	if nombre == "" {
		return Producto{}, fmt.Errorf("El producto debe tener un nombre")
	}
	return Producto{
		Nombre:   nombre,
		Precio:   precio,
		Alergias: alergias,
	}, nil
}