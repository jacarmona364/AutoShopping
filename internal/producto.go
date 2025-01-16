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


type Alergia struct {
	Nombre string
}

func NuevaAlergia(nombre string) (Alergia, error) {
	if nombre == "" {
		return Alergia{}, fmt.Errorf("La alergia debe tener nombre")
	}
	return Alergia{Nombre: nombre}, nil
}


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