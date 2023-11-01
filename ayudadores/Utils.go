package ayudadores

import (
	"image"
	"os"
	_ "image/png"
	"github.com/faiface/pixel"
)

type Posiciones struct {
	posiciones [20]pixel.Vec
}

type ElementoVista struct {
	Identificador int
	sprite pixel.Sprite
	coordenadas pixel.Vec
	Estado bool
}

func NuevoElementoVista(Identificador int, sprite pixel.Sprite, coordenadas pixel.Vec, Estado bool) *ElementoVista {
	return &ElementoVista{
		Identificador: Identificador,
		sprite: sprite,
		coordenadas: coordenadas,
		Estado: Estado,
	}
}

func (EV *ElementoVista) GetData() (pixel.Sprite, pixel.Vec) {
	return	EV.sprite, EV.coordenadas
}

func CargarImagen(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func NuevoSprite(img pixel.Picture) (pixel.Sprite) {
	return *pixel.NewSprite(img, img.Bounds())
}

func NuevasPosiciones() *Posiciones {
	return &Posiciones{
		posiciones: [20]pixel.Vec{
			pixel.V(145, 90),pixel.V(145, 165),pixel.V(145, 240),pixel.V(145, 310),pixel.V(145, 380),
			pixel.V(145, 450),pixel.V(145, 520),pixel.V(145, 590),pixel.V(145, 670),pixel.V(145, 740),

			pixel.V(270, 90),pixel.V(270, 165),pixel.V(270, 240),pixel.V(270, 310),pixel.V(270, 380),
			pixel.V(270, 450),pixel.V(270, 520),pixel.V(270, 590),pixel.V(270, 670),pixel.V(270, 740),
		},
	}
}

func ObtenerCoordenadas(posiciones *Posiciones, posicion int) pixel.Vec {
	return posiciones.posiciones[posicion]
}