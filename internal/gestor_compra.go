package internal

import "errors"

type GestorCompra struct {
	Usuario Usuario
}

func NuevoGestorCompra(usuario Usuario) GestorCompra {
	return GestorCompra{
		Usuario: usuario,
	}
}

func (g *GestorCompra) esProductoCompatible(producto Producto) bool {
	if !g.esSupermercadoValido(producto.Supermercado) {
		return false
	}

	for _, restriccion := range g.Usuario.Alergias {
		for _, apto := range producto.AptoPara {
			if restriccion == apto {
				return false
			}
		}
	}

	return true
}

func (g *GestorCompra) esSupermercadoValido(supermercado Supermercado) bool {
	for _, preferido := range g.Usuario.Supermercados {
		if supermercado == preferido {
			return true
		}
	}
	return false
}
