package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	worldWidth  = 1080
	worldHeight = 1080	
)

type RGBA struct {
	R, G, B, A uint8
}
var (
	redColor   = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	greenColor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	blueColor  = color.RGBA{R: 0, G: 0, B: 255, A: 255}
)

func calculateColorDifference(a, b color.RGBA) float64 {
	redColorA := int(a.R)
	greenColorA := int(a.G)
	blueColorA := int(a.B)

	redColorB := int(b.R)
	greenColorB := int(b.G)
	blueColorB := int(b.B)

	redDifference := redColorA - redColorB
	greenDifference := greenColorA - greenColorB
	blueDifference := blueColorA - blueColorB

	return math.Abs(float64(redDifference)) +
		math.Abs(float64(greenDifference)) +
		math.Abs(float64(blueDifference))
}

func main() {
	ebiten.SetWindowSize(worldWidth, worldHeight)
	ebiten.SetWindowTitle("Simland-Go")

	particles := createRedParticles(5)

	particles = append(particles, createGreenParticles(10)...)
	particles = append(particles, createBlueParticles(0)...)

	game := &Game{
		particles: particles,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
