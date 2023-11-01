package modelos

type Entrada struct {
	estados [3]string
	estadoActual string
}

func NuevaEntrada() *Entrada {
	return &Entrada{
		estados: [3]string{"ENTRANDO", "SALIENDO", "INACTIVO"},
		estadoActual: "INACTIVO",
	}
}

func (E *Entrada) ObtenerEstadoActual() string {
	return E.estadoActual	
}

func (E *Entrada) CambiarEstadoActual(opcion int) {
	E.estadoActual = E.estados[opcion]
}

func (E *Entrada) CambiarEstado(canal chan int) {
	for {
		select {
		case val := <-canal:
			E.CambiarEstadoActual(val)
		}
	}
}