package internal

import (
	"errors"
	"fmt"
)

type Alergia struct {
	Nombre string
}

func NuevaAlergia(nombre string) (Alergia, error) {
	if nombre == "" {
		return Alergia{}, fmt.Errorf("La alergia debe tener nombre")
	}
	return Alergia{Nombre: nombre}, nil
}