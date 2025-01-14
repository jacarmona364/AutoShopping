package internal

import (
	"errors"
	"fmt"
)

type ListaCompra struct {
	Productos []Producto
	Total     Precio
}

func NuevaListaCompra() ListaCompra {
	return ListaCompra{
		Productos: []Producto{}, 
		Total:     Precio{},   
	}
}