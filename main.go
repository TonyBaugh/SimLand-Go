package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	worldWidth  = 1080
	worldHeight = 1080
)

type Game struct {	
	particles []Particle
	RGBA color.RGBA
	debugColorDifference float64
}

type Particle struct {
	xPos float32
	yPos float32
	xVel float32
	yVel float32
	radius float32
	particleColor color.RGBA
}

type RGBA struct {
	R, G, B, A uint8
}

func (p *Particle) Update() {
	p.xPos += p.xVel
	p.yPos += p.yVel

	//Wall Collision Detection
	//Right Wall
	if p.xPos + p.radius >= worldWidth {
		p.xPos = worldWidth - p.radius
		p.xVel *= -1
	}
	// Left Wall
	if p.xPos <= p.radius {
		p.xPos = p.radius
		p.xVel *= -1
	}
	//Bottom wall
	if p.yPos + p.radius >= worldHeight {
		p.yPos = worldHeight - p.radius
		p.yVel *= -1
	}
	//Top wall
	if p.yPos <= p.radius {
		p.yPos = p.radius
		p.yVel *= -1
	}
}

func createParticles(count int) []Particle {
	particleList := make([]Particle, 0, count)
	radius := 3
		
	for range count {
		particle := Particle {
			xPos: float32(rand.IntN(worldWidth - 2*radius) + radius),
			yPos: float32(rand.IntN(worldHeight - 2*radius)+ radius),
			xVel: float32(rand.IntN(10) + 1),
			yVel: float32(rand.IntN(10) + 1),
			radius: float32(radius),
			particleColor: color.RGBA{uint8(rand.IntN(256)), uint8(rand.IntN(256)), uint8(rand.IntN(256)), 255},
		}
		particleList = append(particleList, particle)
	}
	return particleList
}

func calculateColorDifference(a, b color.RGBA) (float64) {
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

func (g *Game) Update() error {
	g.debugColorDifference = calculateColorDifference(g.particles[0].particleColor, g.particles[1].particleColor)
	for i := range g.particles{
		g.particles[i].Update()		
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.particles {
	vector.FillCircle(screen, g.particles[i].xPos, g.particles[i].yPos, g.particles[i].radius, g.particles[i].particleColor, true)
	}
	ebitenutil.DebugPrint(
	screen,
	fmt.Sprintf("Particle 0: R %d, G %d, B %d\nParticle 1: R %d, G %d, B %d\nColor difference: %.0f", 
	g.particles[0].particleColor.R, g.particles[0].particleColor.G, g.particles[0].particleColor.B,
	g.particles[1].particleColor.R, g.particles[1].particleColor.G, g.particles[1].particleColor.B,
	g.debugColorDifference),
	
)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return worldWidth, worldHeight
}

func main() {
	ebiten.SetWindowSize(worldWidth, worldHeight)
	ebiten.SetWindowTitle("Simland-Go")

	
	game := &Game{
		particles: createParticles(2),
		}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
