package main

import (
	"fmt"
	"image/color"
	
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	particles            []Particle
	RGBA                 color.RGBA
	debugColorDifference float64
}

func (g *Game) FindClosestParticle(hunterIndex int, targetColor color.RGBA) (targetIndex int, targetDistance float64, offsetX int, offsetY int) {
	closestTarget := -1
	closestTargetDistance := worldWidth * 2.0
	targetOffsetX, targetOffsetY := 0, 0
	// For every particle in the list
	for i := range g.particles {
		// If the particle we're looking at is the target color sent into the function
		if g.particles[i].particleColor == targetColor {
			// Calculate offset and distance
			offsetX, offsetY := particleOffset(int(g.particles[hunterIndex].xPos), int(g.particles[i].xPos), int(g.particles[hunterIndex].yPos), int(g.particles[i].yPos))
			distance := particleDistance(offsetX, offsetY)
			// Update closest target and its distance
			if distance < closestTargetDistance {
				closestTargetDistance = distance
				closestTarget = i
				targetOffsetX = offsetX
				targetOffsetY = offsetY
			}
		}
	}
	// Return closest target and its distance
	return closestTarget, closestTargetDistance, targetOffsetX, targetOffsetY
}
func (g *Game) Update() error {
	g.debugColorDifference = calculateColorDifference(g.particles[0].particleColor, g.particles[1].particleColor)

	// blueParticleIndex := len(g.particles) - 1
	// redXPos, redYPos := g.particles[0].xPos, g.particles[0].yPos
	// blueXPos, blueYPos := g.particles[blueParticleIndex].xPos, g.particles[blueParticleIndex].yPos

	var (		
		eatenGreen = -1				
	)	
	var greensToBeAdded []Particle
	
	// Find closest green to red dot
	for i := range g.particles {
		if g.particles[i].particleColor == greenColor {
				g.particles[i].distanceTraveled += math.Abs(float64(g.particles[i].xVel)) + math.Abs(float64(g.particles[i].yVel))
				
				if g.particles[i].distanceTraveled > 500 {
					greensToBeAdded = append(greensToBeAdded, createGreenParticles(1)...)
					g.particles[i].distanceTraveled = 0
				}
			}
		if g.particles[i].particleColor == redColor {
			closestGreenToRed, nearestGreenDistanceToRed, offsetX, offsetY :=
				g.FindClosestParticle(i, greenColor)

			if closestGreenToRed == -1 {
				continue
			}
			
			
			//  Check if red is close enough to eat green
			if nearestGreenDistanceToRed < 6 {									
				eatenGreen = closestGreenToRed
				break
			}

			if offsetX > 3 && nearestGreenDistanceToRed > 6 {
				g.particles[i].xVel = -5
			} else if offsetX < -3 && nearestGreenDistanceToRed > 6 {
				g.particles[i].xVel = 5
			} else {
				g.particles[i].xVel = 0
			}

			if offsetY > 3 && nearestGreenDistanceToRed > 6 {
				g.particles[i].yVel = -5
			} else if offsetY < -3 && nearestGreenDistanceToRed > 6 {
				g.particles[i].yVel = 5
			} else {
				g.particles[i].yVel = 0
			}

		}
	
	}
	// Rebuild list with eatenGreen cut out of it, preserving order.
	if eatenGreen != -1 { 
		g.particles = append(g.particles[0:eatenGreen], g.particles[eatenGreen+1:]... )
	}
	g.particles = append(g.particles, greensToBeAdded...)
	

	// red follows green
	// if xdiff is positive, red is to the right of green, red needs to move negative(left)

	// if grnRedXDiff > 3 && grnRedDistance > 6 {
	// 	g.particles[0].xVel = -3
	// } else if grnRedXDiff < -3 && grnRedDistance > 6 {
	// 	g.particles[0].xVel = 3
	// } else {
	// 	g.particles[0].xVel = 0
	// }

	// if grnRedYDiff > 3 && grnRedDistance > 6 {
	// 	g.particles[0].yVel = -3
	// } else if grnRedYDiff < -3 && grnRedDistance > 6 {
	// 	g.particles[0].yVel = 3
	// } else {
	// 	g.particles[0].yVel = 0
	// }

	// Blue is scared of red
	// if redBlueDistance < 150 {
	// 	if redBlueXDiff > 5 {
	// 		g.particles[blueParticleIndex].xVel = -5
	// 	} else if redBlueXDiff < -5 {
	// 		g.particles[blueParticleIndex].xVel = 5
	// 	} else {
	// 		g.particles[blueParticleIndex].xVel = 0
	// 	}
	// 	if redBlueYDiff > 5 {
	// 		g.particles[blueParticleIndex].yVel = -5
	// 	} else if redBlueYDiff < -5 {
	// 		g.particles[blueParticleIndex].yVel = 5
	// 	} else {
	// 		g.particles[blueParticleIndex].yVel = 0
	// 	}

	// } else {
	// 	// Blue follows green
	// 	if grnBlueXDiff > 4 && grnBlueDistance > 6 {
	// 		g.particles[blueParticleIndex].xVel = -4
	// 	} else if grnBlueXDiff < -4 && grnBlueDistance > 6 {
	// 		g.particles[blueParticleIndex].xVel = 4
	// 	} else {
	// 		g.particles[blueParticleIndex].xVel = 0
	// 	}

	// 	if grnBlueYDiff > 4 && grnBlueDistance > 6 {
	// 		g.particles[blueParticleIndex].yVel = -4
	// 	} else if grnBlueYDiff < -4 && grnBlueDistance > 6 {
	// 		g.particles[blueParticleIndex].yVel = 4
	// 	} else {
	// 		g.particles[blueParticleIndex].yVel = 0
	// 	}
	//}

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
