package internal

import (
	"errors"
)

type GestorCompra struct {
	productosDisponibles   map[string][]models.Producto
	preferenciaUsuario     models.PreferenciaUsuario
}

func NewGestorCompra(productos []models.Producto, preferencia models.PreferenciaUsuario) (*GestorCompra, error) {
	if len(productos) == 0 {
		return nil, errors.New("no hay productos disponibles")
	}

	productosMap := make(map[string][]models.Producto)
	for _, p := range productos {
		productosMap[p.Supermercado] = append(productosMap[p.Supermercado], p)
	}

	for _, supermercado := range preferencia.Supermercados {
		if _, existe := productosMap[supermercado]; !existe {
			return nil, errors.New("el supermercado " + supermercado + " no está disponible")
		}
	}

	return &GestorCompra{
		productosDisponibles: productosMap,
		preferenciaUsuario:   preferencia,
	}, nil
}

func (g *GestorCompra) GenerarListaCompra() (*models.ListaCompra, error) {
	var lista models.ListaCompra

	for _, supermercado := range g.preferenciaUsuario.Supermercados {
		for _, producto := range g.productosDisponibles[supermercado] {
			if validaciones.EsProductoValido(producto, g.preferenciaUsuario.Alergias) && 
			   lista.Total+producto.Precio <= g.preferenciaUsuario.Presupuesto {
				lista.Productos = append(lista.Productos, producto)
				lista.Total += producto.Precio
			}
		}
	}

	if len(lista.Productos) == 0 {
		return nil, errors.New("no se encontraron productos válidos")
	}
	return &lista, nil
}
