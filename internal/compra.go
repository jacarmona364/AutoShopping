package internal

import (
	"errors"
	"fmt"
)

type ListaCompra struct {
	Productos []Producto
	Total     float64
}

func esProductoValido(producto Producto, alergias []Alergia) bool {
	for _, alergen := range alergias {
		for _, apto := range producto.Alergias {
			if alergen == apto {
				return false
			}
		}
	}
	return true
}

type Supermercado string

const (
	Mercadona		Supermercado = "Mercadona"
	Alcampo			Supermercado = "Alcampo"
	Carrefour		Supermercado = "Carrefour"
	CorteIngles		Supermercado = "Corte Ingles"
	Aldi			Supermercado = "Aldi"
	Dia				Supermercado = "Dia"
	Lidl			Supermercado = "Lidl"
	Coviran			Supermercado = "Coviran"
)