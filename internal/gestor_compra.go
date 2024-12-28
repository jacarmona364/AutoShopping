package internal

import "errors"

type GestorCompra struct {
	Usuario Usuario
	ListaCompra ListaCompra
}

func NuevoGestorCompra(usuario Usuario, listaCompra ListaCompra) GestorCompra {
	return GestorCompra{
		Usuario: usuario,
		ListaCompra: listaCompra,
	}
}
