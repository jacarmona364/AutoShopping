package internal

import (
	"errors"
	"fmt"
)

type ListaCompra struct {
	Productos []Producto
	Total     Precio
}

func NuevaListaCompra(productos []Producto) (ListaCompra, error) {
    if len(productos) == 0 {
		return ListaCompra{}, fmt.Errorf("la lista de productos no puede estar vacía")
	}

	var total float64
	for _, producto := range productos {
		total += producto.Precio.Valor
	}

	totalPrecio, err := NuevoPrecio(total)
	if err != nil {
		return ListaCompra{}, fmt.Errorf("error al calcular el total: %w", err)
	}

	return ListaCompra{
		Productos: productos,
		Total:     totalPrecio,
	}, nil
}