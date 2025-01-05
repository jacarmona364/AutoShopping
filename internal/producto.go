package internal

import "errors"

type Producto struct {
	Nombre       string
	Marca        string
	Precio       float64
	Supermercado Supermercado
	Alergias     []Alergia
}

func NuevoProducto(nombre, marca string, precio float64, supermercado Supermercado, alergias []Alergia) Producto {
	if precio < 0 {
        return Producto{}, errors.New("El precio del producto no puede ser negativo")
    }

	return Producto{
		Nombre:       nombre,
		Marca:        marca,
		Precio:       precio,
		Supermercado: supermercado,
		Alergias:     alergias,
	}
}

type Alergia string

const (
	SinGluten       Alergia = "SinGluten"
	SinLactosa      Alergia = "SinLactosa"
	Vegetariano     Alergia = "Vegetariano"
	Vegano          Alergia = "Vegano"
	SinFrutosSecos  Alergia = "SinFrutosSecos"
	SinHuevos       Alergia = "SinHuevos"
	SinMariscos     Alergia = "SinMariscos"
	SinSoja         Alergia = "SinSoja"
	SinPescado      Alergia = "SinPescado"
)