package internal

import (
	"errors"
	"fmt"
)

type Producto struct {
	Nombre      string
	Categoria   string
	Precio      float64
	AptoPara    []string  
	Supermercado string
}

type PreferenciaUsuario struct {
	Alergias      []string
	Presupuesto   float64
	Supermercados []string
}

type ListaCompra struct {
	Productos []Producto
	Total     float64
}

func esProductoValido(producto Producto, alergias []string) bool {
	for _, alergen := range alergias {
		for _, apto := range producto.AptoPara {
			if alergen == apto {
				return false
			}
		}
	}
	return true
}
