package internal

import "errors"

type Producto struct {
	Nombre       string
	Marca        string
	Precio       float64
	Supermercado Supermercado
	AptoPara     []TipoRestriccion //Etiquetas para alergias o restricciones
}

func NuevoProducto(nombre, marca string, precio float64, supermercado string, aptoPara []string) Producto {
	return Producto{
		Nombre:       nombre,
		Marca:        marca,
		Precio:       precio,
		Supermercado: supermercado,
		AptoPara:     aptoPara,
	}
}

type TipoRestriccion string

const (
	SinGluten       TipoRestriccion = "SinGluten"
	SinLactosa      TipoRestriccion = "SinLactosa"
	Vegetariano     TipoRestriccion = "Vegetariano"
	Vegano          TipoRestriccion = "Vegano"
	SinFrutosSecos  TipoRestriccion = "SinFrutosSecos"
	SinHuevos       TipoRestriccion = "SinHuevos"
	SinMariscos     TipoRestriccion = "SinMariscos"
	SinSoja         TipoRestriccion = "SinSoja"
	SinPescado      TipoRestriccion = "SinPescado"
)