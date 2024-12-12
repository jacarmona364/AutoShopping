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

type GestorCompra struct {
	productosDisponibles   map[string][]Producto
	preferenciaUsuario     PreferenciaUsuario
}

func NewGestorCompra(productos []Producto, preferencia PreferenciaUsuario) (*GestorCompra, error) {
	if len(productos) == 0 {
		return nil, errors.New("debe haber al menos un producto disponible")
	}

	productosMap := make(map[string][]Producto)
	for _, p := range productos {
		productosMap[p.Supermercado] = append(productosMap[p.Supermercado], p)
	}

	for _, supermercado := range preferencia.Supermercados {
		if _, existe := productosMap[supermercado]; !existe {
			return nil, fmt.Errorf("el supermercado %s no tiene productos registrados", supermercado)
		}
	}

	return &GestorCompra{
		productosDisponibles: productosMap,
		preferenciaUsuario:   preferencia,
	}, nil
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

func (g *GestorCompra) GenerarListaCompra() (*ListaCompra, error) {
	var lista ListaCompra

	for _, supermercado := range g.preferenciaUsuario.Supermercados {
		for _, producto := range g.productosDisponibles[supermercado] {
			if esProductoValido(producto, g.preferenciaUsuario.Alergias) {
				if lista.Total+producto.Precio <= g.preferenciaUsuario.Presupuesto {
					lista.Productos = append(lista.Productos, producto)
					lista.Total += producto.Precio
				}
			}
		}
	}

	if len(lista.Productos) == 0 {
		return nil, errors.New("no se encontraron productos válidos dentro del presupuesto y preferencias")
	}

	return &lista, nil
}
