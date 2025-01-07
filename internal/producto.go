package internal

import "errors"

type Producto struct {
	Nombre       string
	Marca        string
	Precio       float64
	Supermercado Supermercado
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

type Supermercado string const (
	Mercadona		Supermercado = "Mercadona"
	Alcampo			Supermercado = "Alcampo"
	Carrefour		Supermercado = "Carrefour"
	CorteIngles		Supermercado = "Corte Ingles"
	Aldi			Supermercado = "Aldi"
	Dia				Supermercado = "Dia"
	Lidl			Supermercado = "Lidl"
	Coviran			Supermercado = "Coviran"
)

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