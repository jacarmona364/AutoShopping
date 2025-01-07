package internal

import (
	"errors"
	"fmt"
)

type ListaCompra struct {
	Productos []Producto
	Total     float64
}

func NuevaListaCompra(productos []Producto) ListaCompra {
    var total float64
    for _, producto := range productos {
        total += producto.Precio
    }
    return ListaCompra{
        Productos: productos,
        Total:     total,
    }
}