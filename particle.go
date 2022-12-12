package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Particle struct {
	RayCaster
	pos  pixel.Vec
	rays []*Ray
}

// NewParticle returns a Particle at the specified position vector pos.
func (rayCaster *RayCaster) NewParticle(pos pixel.Vec) *Particle {
	var rays []*Ray
	var i float64

	// appending rays at each angle to the particle
	for i = 0; i < 360; i += 1 {
		var radian float64 = i * 0.01745
		ray := rayCaster.NewRay(pos, radian)
		rays = append(rays, ray)
	}

	return &Particle{
		RayCaster: *rayCaster,
		pos:       pos,
		rays:      rays,
	}
}

// update updates the position of the particle by the offset x in horizontal and y in vertical direction.
func (particle *Particle) update(x, y float64) {
	particle.pos.X += x
	particle.pos.Y += y
	particle.updateRays()
}

// updateRays refreshes the rays of the Particle after it has been moved
func (particle *Particle) updateRays() {
	for _, ray := range particle.rays {
		ray.pos = particle.pos
	}
}

// look casts each Ray of the Particle to each Boundary and renders the closest intersecting Ray
func (particle *Particle) look() {
	for i := 0; i < len(particle.rays); i++ {
		ray := particle.rays[i]

		record := math.MaxFloat64

		var closest *pixel.Vec
		closest = nil

		for _, wall := range particle.RayCaster.boundaries {
			point := ray.cast(&wall)
			if point != nil {
				d := dist(particle.pos, *point)

				if d < record {
					record = d
					closest = point
				}
			}
		}

		if closest != nil {
			var rayColor color.RGBA
			rayColor.R = 1
			rayColor.G = 1
			rayColor.B = 1
			rayColor.A = 0
			particle.RayCaster.imd.Color = rayColor
			particle.RayCaster.imd.EndShape = imdraw.SharpEndShape
			particle.RayCaster.imd.Push(particle.pos)
			particle.RayCaster.imd.Push(*closest)
			particle.RayCaster.imd.Line(1)
		}
	}
}

// dist finds out the distance between 2 vectors
func dist(v1, v2 pixel.Vec) float64 {
	x := v1.X - v2.X
	y := v1.Y - v2.Y
	return math.Abs(x + y)
}

// show renders the particle
func (particle *Particle) show() {
	var rayColor color.RGBA
	rayColor.R = 1
	rayColor.G = 1
	rayColor.B = 1
	rayColor.A = 0
	particle.RayCaster.imd.Color = rayColor
	particle.RayCaster.imd.Push(particle.pos)
	particle.RayCaster.imd.Ellipse(pixel.V(10, 10), 0)
}
