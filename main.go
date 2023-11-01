package main

import (
	"Estacionamiento-go/ayudadores"
	"Estacionamiento-go/controladores"
	"Estacionamiento-go/modelos"
	_ "fmt"
	_ "image/png"
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 422, 800),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	estacionamientoImagen, err := ayudadores.CargarImagen("./assets/Estacionamiento.png")
	if err != nil {
		panic(err)
	}
	estacionamientoSprite := ayudadores.NuevoSprite(estacionamientoImagen)

	// Canales de Tipos
	canalAuto := make(chan modelos.Auto)

	// Semaforo
	mu := &sync.Mutex{}

	// Semaforo auxiliat
	canalAuxEntrada := make(chan int)
	canalVista := make(chan ayudadores.ElementoVista)

	// Controladores
	controladorEstacionamiento := controladores.NuevoEstacionamientoControlador(mu)
	controladorAutos := controladores.NuevoAutoControlador()
	controladorEntrada := controladores.NuevoControladorEntrada()

	// Aqui estara escuchando los autos generados que vienen del auto del canal
	go controladorEstacionamiento.EstacionarAutos(canalAuto, controladorEntrada, canalAuxEntrada, canalVista)
	go controladorAutos.GenerarAutos(canalAuto)
	go controladorEntrada.Entrada.CambiarEstado(canalAuxEntrada)

	// Aqui empezara a generar autos y enviarlos por el canal de los autos

	var arreglo []ayudadores.ElementoVista
	for !win.Closed() {

		estacionamientoSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		select {
		case val, ok := <-canalVista:
			if !ok {
				return
			}
			if val.Estado {
				arreglo = append(arreglo, val)
			} else {
				var arregloAux []ayudadores.ElementoVista
				for _, elemento := range arreglo {
					if elemento.Identificador != val.Identificador {
						arregloAux = append(arregloAux, elemento)
					}
				}
				arreglo = arreglo[:0]
				arreglo = append(arreglo, arregloAux...)
			}
		}

		for _, elemento := range arreglo {
			sprite, coo := elemento.GetData()
			sprite.Draw(win, pixel.IM.Moved(coo))
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

// Para optimizar se uso Batch

// Requisito; Crear un batch := pixel.NewBatch(&pixel.TrianglesData{}, autoImagen) //el autoImagen es el auto que se guardara
// Primero consiste en crear un arreglo de sprites que almacene pixel.Sprite
// Segundo Guardar los sprites en el arreglo
// Tercero Recorremos el arreglo y usaremos el draw de cada elemento sprite.Draw(batch, pixel.IM.Moved(pixel.V(270, float64(90 + i*75))))
// Cuarto Mandamos a pintar todo con batch.Draw(win)
