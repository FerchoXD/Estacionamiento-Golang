package modelos

import (
	"Estacionamiento-go/ayudadores"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/faiface/pixel"
)

type Auto struct {
	Identificador int
	Tiempo        int
}

func NuevoAuto() *Auto {
	rand.Seed(time.Now().UnixNano())
	tiempo := rand.Intn(5) + 10
	return &Auto{
		Tiempo: tiempo,
	}
}

func (a *Auto) GenerarAutos(canal chan Auto) {
	for i := 1; i <= 100; i++ {
		car := NuevoAuto()
		car.PonerIdentificador(i)
		canal <- *car
		rand.Seed(time.Now().UnixNano())
		tiempo := rand.Intn(2) + 1
		time.Sleep(time.Second * time.Duration(tiempo))
	}
	close(canal)
}

func (a *Auto) PonerIdentificador(Identificador int) {
	a.Identificador = Identificador
}

func (a *Auto) Estacionar(estacionamientoControler *Estacionamiento, mu *sync.Mutex, canalEntrada chan int, posicion int, coordenadas pixel.Vec, ch chan ayudadores.ElementoVista) {
	mu.Lock()
	estacionamientoControler.nEspaciosDisponibles--
	fmt.Println("El auto", a, "se acaba de estacionar")
	fmt.Println("Quedan solo", estacionamientoControler.nEspaciosDisponibles, "disponibles")

	imagen, _ := ayudadores.CargarImagen("./assets/auto.png")
	elemento := ayudadores.NuevoElementoVista(a.Identificador, ayudadores.NuevoSprite(imagen), coordenadas, true)
	ch <- *elemento

	canalEntrada <- 0
	mu.Unlock()

	time.Sleep(time.Second * time.Duration(a.Tiempo))

	mu.Lock()

	estacionamientoControler.nEspaciosDisponibles++
	estacionamientoControler.espaciosDisponibles[posicion] = true

	imagen, _ = ayudadores.CargarImagen("./assets/auto.png")
	elemento = ayudadores.NuevoElementoVista(a.Identificador, ayudadores.NuevoSprite(imagen), coordenadas, false)
	ch <- *elemento

	fmt.Println("El auto", a, "se acaba de ir")
	fmt.Println("Quedan solo", estacionamientoControler.nEspaciosDisponibles, "disponibles")
	canalEntrada <- 1
	mu.Unlock()
}
