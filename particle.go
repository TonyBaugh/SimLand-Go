package main

import (
	"image/color"
	"math"
	"math/rand/v2"
)

type Particle struct {
	xPos          float32
	yPos          float32
	xVel          float32
	yVel          float32
	radius        float32
	particleColor color.RGBA
}

const (
	maxXVel = 10
	maxYVel = 10
)

func (p *Particle) Update() {
	p.xPos += p.xVel
	p.yPos += p.yVel

	//Wall Collision Detection
	//Right Wall
	if p.xPos+p.radius >= worldWidth {
		p.xPos = worldWidth - p.radius
		p.xVel *= -1
	}
	// Left Wall
	if p.xPos <= p.radius {
		p.xPos = p.radius
		p.xVel *= -1
	}
	//Bottom wall
	if p.yPos+p.radius >= worldHeight {
		p.yPos = worldHeight - p.radius
		p.yVel *= -1
	}
	//Top wall
	if p.yPos <= p.radius {
		p.yPos = p.radius
		p.yVel *= -1
	}
}

func particleOffset(p1XDist, p2XDist, p1YDist, p2YDist int) (xDiff int, yDiff int) {
	xDiff = p1XDist - p2XDist
	yDiff = p1YDist - p2YDist

	return xDiff, yDiff
}

func particleDistance(xDiff, yDiff int) (distance float64) {
	//xdiff*xdiff + ydiff*ydiff = distance*2
	return math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff))
}

func createParticles(count int) []Particle {
	particleList := make([]Particle, 0, count)
	radius := 3

	for range count {
		particle := Particle{
			xPos:          float32(rand.IntN(worldWidth-2*radius) + radius),
			yPos:          float32(rand.IntN(worldHeight-2*radius) + radius),
			xVel:          float32(rand.IntN(maxXVel) + 1), // Random velocity, (range max) + {Shift range upwards}
			yVel:          float32(rand.IntN(maxYVel) + 1),
			radius:        float32(radius),
			particleColor: color.RGBA{uint8(rand.IntN(256)), uint8(rand.IntN(256)), uint8(rand.IntN(256)), 255},
		}
		particleList = append(particleList, particle)
	}
	return particleList
}

func createRedParticles(count int) []Particle {
	particleList := make([]Particle, 0, count)
	radius := 3

	for range count {
		particle := Particle{
			xPos:          float32(rand.IntN(worldWidth-2*radius) + radius),
			yPos:          float32(rand.IntN(worldHeight-2*radius) + radius),
			xVel:          float32(3),
			yVel:          float32(3),
			radius:        float32(radius),
			particleColor: color.RGBA{255, 0, 0, 255},
		}
		particleList = append(particleList, particle)
	}
	return particleList
}
func createGreenParticles(count int) []Particle {
	particleList := make([]Particle, 0, count)
	radius := 3

	for range count {
		particle := Particle{
			xPos:          float32(rand.IntN(worldWidth-2*radius) + radius),
			yPos:          float32(rand.IntN(worldHeight-2*radius) + radius),
			xVel:          float32(5),
			yVel:          float32(5),
			radius:        float32(radius),
			particleColor: color.RGBA{0, 255, 0, 255},
		}
		particleList = append(particleList, particle)
	}
	return particleList
}
func createBlueParticles(count int) []Particle {
	particleList := make([]Particle, 0, count)
	radius := 3

	for range count {
		particle := Particle{
			xPos:          float32(rand.IntN(worldWidth-2*radius) + radius),
			yPos:          float32(rand.IntN(worldHeight-2*radius) + radius),
			xVel:          float32(3),
			yVel:          float32(3),
			radius:        float32(radius),
			particleColor: color.RGBA{0, 0, 255, 255},
		}
		particleList = append(particleList, particle)
	}
	return particleList
}
