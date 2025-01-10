package internal

import "errors"

type Producto struct {
	Nombre       string
	Marca        string
	Precio       float64
	Supermercado float64
	Alergias     []Alergia
}

type Alergia string const (
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


func NuevoProducto(nombre, marca string, precio float64, supermercado float64, alergias []Alergia) Producto {
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