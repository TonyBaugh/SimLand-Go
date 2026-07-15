package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	worldWidth  = 1080
	worldHeight = 1080
)

type Game struct {
	particles            []Particle
	RGBA                 color.RGBA
	debugColorDifference float64
}

type RGBA struct {
	R, G, B, A uint8
}

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

func (g *Game) Update() error {
	g.debugColorDifference = calculateColorDifference(g.particles[0].particleColor, g.particles[1].particleColor)	
	
	
	blueParticleIndex := len(g.particles)-1
	redXPos, redYPos := g.particles[0].xPos, g.particles[0].yPos
	blueXPos, blueYPos := g.particles[blueParticleIndex].xPos, g.particles[blueParticleIndex].yPos
	// make red track nearest green
	var nearestGreenDistance = worldHeight*2.0 // bigger than possible
	var closestGreen int
	for i := range g.particles {
		if g.particles[i].particleColor.G == 255 {
			offsetX, offsetY := particleOffset(int(redXPos), int(g.particles[i].xPos), int(redYPos), int(g.particles[i].yPos))
			distance := particleDistance(offsetX, offsetY)
			if distance < nearestGreenDistance {
				nearestGreenDistance = distance
				closestGreen = i
			}
		}		
	}
	//log.Printf("xdiff %d, yDiff %d", xDiff, yDiff)

	
	grnRedXDiff, grnRedYDiff := particleOffset(int(g.particles[0].xPos), int(g.particles[closestGreen].xPos), int(g.particles[0].yPos), int(g.particles[closestGreen].yPos))
	grnRedDistance := particleDistance(grnRedXDiff, grnRedYDiff)

	
	redBlueXDiff, redBlueYDiff := particleOffset(int(redXPos), int(blueXPos), int(redYPos), int(blueYPos))
	redBlueDistance := particleDistance(redBlueXDiff, redBlueYDiff)

	grnBlueXDiff, grnBlueYDiff := particleOffset(int(g.particles[blueParticleIndex].xPos), int(g.particles[closestGreen].xPos), int(g.particles[blueParticleIndex].yPos), int(g.particles[closestGreen].yPos))
	grnBlueDistance := particleDistance(grnBlueXDiff, grnBlueYDiff)
	//log.Printf("blue/red distance: %.0f", redBlueDistance)
	
	
	// Blue is scared of red
	if redBlueDistance < 150 {
		if redBlueXDiff > 5 {
			g.particles[blueParticleIndex].xVel = -5
		} else if redBlueXDiff < -5 {
			g.particles[blueParticleIndex].xVel = 5
		} else if redBlueYDiff > 5 {
			g.particles[blueParticleIndex].yVel = -5	
		} else if redBlueYDiff < -5 {
			g.particles[blueParticleIndex].yVel = 5	
		}
	}
	
	// Blue follows green
	if grnBlueXDiff > 4 && grnBlueDistance > 6{
		g.particles[blueParticleIndex].xVel = -4
	} else if grnBlueXDiff < -4 && grnBlueDistance > 6{
		g.particles[blueParticleIndex].xVel = 4
	} else {
		g.particles[blueParticleIndex].xVel = 0
	}

	if grnBlueYDiff > 4 && grnBlueDistance > 6{
		g.particles[blueParticleIndex].yVel = -4
	} else if grnBlueYDiff < -4 && grnBlueDistance > 6{
		g.particles[blueParticleIndex].yVel = 4
	} else {
		g.particles[blueParticleIndex].yVel = 0
	}

	// red follows green
	//p[0] = red, p[1] = green
	// if xdiff is positive, red is to the right of green, red needs to move negative(left)
	if grnRedXDiff > 3 && grnRedDistance > 6{
		g.particles[0].xVel = -3
	} else if grnRedXDiff < -3 && grnRedDistance > 6{
		g.particles[0].xVel = 3
	} else {
		g.particles[0].xVel = 0
	}
	
	if grnRedYDiff > 3 && grnRedDistance > 6{
		g.particles[0].yVel = -3
	} else if grnRedYDiff < -3 && grnRedDistance > 6{
		g.particles[0].yVel = 3
	} else {
		g.particles[0].yVel = 0
	}


	for i := range g.particles {
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

	particles := createRedParticles(1)

	particles = append(particles, createGreenParticles(20)...)
	particles = append(particles, createBlueParticles(1)...)

	game := &Game{
		particles: particles,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
