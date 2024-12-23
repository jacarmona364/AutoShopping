package internal

type Usuario struct {
	DNI           string
	Nombre        string
	Presupuesto   float64
	Alergias      []TipoRestriccion
	Supermercados []Supermercado
}

