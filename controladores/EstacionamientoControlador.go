package controladores

import (
	"Estacionamiento-go/ayudadores"
	"Estacionamiento-go/modelos"
	"sync"
)

type EstacionamientoControlador struct {
	estacionamiento modelos.Estacionamiento
	mu              *sync.Mutex
}

func NuevoEstacionamientoControlador(mu *sync.Mutex) *EstacionamientoControlador {
	return &EstacionamientoControlador{
		estacionamiento: *modelos.NuevoEstacionamiento(),
		mu:              mu,
	}
}

func (es *EstacionamientoControlador) EstacionarAutos(canal chan modelos.Auto, controladorEntrada *EntradaControlador, canalEntrada chan int, canalVista chan ayudadores.ElementoVista) {
	aux := true
	posiciones := ayudadores.NuevasPosiciones()
	for aux {
		select {
		case car, ok := <-canal:
			if !ok {
				aux = false
				break
			}
			posicion := es.estacionamiento.ObtenerEspaciosDisponibles()
			coordenadas := ayudadores.ObtenerCoordenadas(posiciones, posicion)
			if es.estacionamiento.ObtenerEspacios() > 0 {
				if controladorEntrada.Entrada.ObtenerEstadoActual() == "ENTRANDO" || controladorEntrada.Entrada.ObtenerEstadoActual() == "INACTIVO" {
					go car.Estacionar(&es.estacionamiento, es.mu, canalEntrada, posicion, coordenadas, canalVista)
				} else {
					canalEntrada <- 0
					go car.Estacionar(&es.estacionamiento, es.mu, canalEntrada, posicion, coordenadas, canalVista)
				}
			}
		}
	}
}
