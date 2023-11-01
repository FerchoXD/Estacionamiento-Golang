package controladores

import "Estacionamiento-go/modelos"

type AutoControlador struct {
	modelo modelos.Auto
}

func NuevoAutoControlador() *AutoControlador {
	return &AutoControlador{
		modelo: *modelos.NuevoAuto(),
	}
}

func (ac *AutoControlador) GenerarAutos(canal chan modelos.Auto) {
	ac.modelo.GenerarAutos(canal)
}
