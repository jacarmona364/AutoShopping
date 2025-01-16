package internal

import (
	"errors"
	"fmt"
)

type Cliente struct {
	DNI           string
	Presupuesto   Precio
	Alergias      []Alergia
	Supermercados []Supermercado
}

func NuevoCliente(dni string, presupuesto float64, alergias []Alergia, supermercados []Supermercado) (Cliente, error) {
	if dni == "" {
		return Cliente{}, fmt.Errorf("El DNI no puede estar vacío")
	}

	return Cliente{
		DNI:           dni,
		Presupuesto:   presupuesto,
		Alergias:      alergias,
		Supermercados: supermercados,
	}, nil
}