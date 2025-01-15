package internal

import "errors"

type Producto struct {
	Nombre       string
	Marca        string
	Precio       Precio
}

func NuevoProducto(nombre, marca string, precio Precio) Producto {
	return Producto{
		Nombre:       nombre,
		Marca:        marca,
		Precio:       precio,
	}
}