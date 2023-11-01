package controladores

import "Estacionamiento-go/modelos"

type EntradaControlador struct {
	Entrada modelos.Entrada
}

func NuevoControladorEntrada() *EntradaControlador {
	return &EntradaControlador{
		Entrada: *modelos.NuevaEntrada(),
	}
}
