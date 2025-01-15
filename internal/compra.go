package internal

import (
	"errors"
	"fmt"
)

type ListaCompra struct {
	Productos []Producto
}

func NuevaListaCompra() ListaCompra {
	return ListaCompra{
		Productos: []Producto{}, 
	}
}