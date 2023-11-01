package modelos

type Estacionamiento struct {
	nEspaciosDisponibles int
	espaciosDisponibles  [20]bool
}

func NuevoEstacionamiento() *Estacionamiento {
	return &Estacionamiento{
		nEspaciosDisponibles: 20,
		espaciosDisponibles: [20]bool {true, true, true, true, true, true, true, true, true, true,
										true, true, true, true, true, true, true, true, true, true},
	}
}

func (e *Estacionamiento) ObtenerEspacios() int {
	return e.nEspaciosDisponibles
}

func (e *Estacionamiento) Aumentar() {
	e.nEspaciosDisponibles++
}

func (e *Estacionamiento) Decrementar() {
	e.nEspaciosDisponibles--
}

func (e *Estacionamiento) ObtenerEspaciosDisponibles() int {
	for i, disponible := range e.espaciosDisponibles {
        if disponible {
            e.espaciosDisponibles[i] = false
            return i
        }
    }
    return -1
}