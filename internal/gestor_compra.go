package internal

import "errors"

type GestorCompra struct {
	Compra map[string]ListaCompra
}

func NuevoGestorCompra(usuario Usuario, listaCompra ListaCompra) GestorCompra {
	return GestorCompra{
		Compra: make(map[string]ListaCompra),
	}
}
