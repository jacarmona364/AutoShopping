package internal

type Usuario struct {
	DNI           string
	Nombre        string
	Presupuesto   float64
	Alergias      []Alergia
	Supermercados []Supermercado
}

func NuevoUsuario(dni, nombre string, presupuesto float64, alergias []Alergia, supermercados []Supermercado) (Usuario, error) {
    if presupuesto < 0 {
        return Usuario{}, errors.New("El presupuesto no puede ser negativo")
    }
    return Usuario{
        DNI:           dni,
        Nombre:        nombre,
        Presupuesto:   presupuesto,
        Alergias:      alergias,
        Supermercados: supermercados,
    }, nil
}