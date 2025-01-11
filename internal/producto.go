package internal

import "errors"

//Objeto Valor: Supernercado
type Supermercado struct {
	Nombre string
}

func NuevoSupermercado(nombre string) (Supermercado, error) {
	if nombre == "" {
		return Supermercado{}, fmt.Errorf("El supermercado debe tener nombre")
	}
	return Supermercado{Nombre: nombre}, nil
}

//Objeto Valor: Precio
type Precio struct {
	Valor float64
}

func NuevoPrecio(valor float64) (Precio, error) {
	if valor < 0 {
		return Precio{}, fmt.Errorf("El precio no puede ser negativo")
	}
	return Precio{Valor: valor}, nil
}

//Objeto valor: Alergia
type Alergia struct {
	Nombre string
}

func NuevaAlergia(nombre string) (Alergia, error) {
	if nombre == "" {
		return Alergia{}, fmt.Errorf("La alergia debe tener nombre")
	}
	return Alergia{Nombre: nombre}, nil
}

//Entidad: Producto
type Producto struct {
	Nombre       string
	Marca        string
	Precio       Precio
	Supermercado Supermercado
	Alergias     []Alergia
}

func NuevoProducto(nombre, marca string, precio Precio, supermercado Supermercado, alergias []Alergia) Producto {
	return Producto{
		Nombre:       nombre,
		Marca:        marca,
		Precio:       precio,
		Supermercado: supermercado,
		Alergias:     alergias,
	}
}