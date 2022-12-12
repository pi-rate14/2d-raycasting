package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Ray struct {
	RayCaster
	pos pixel.Vec
	dir pixel.Vec
}

// NewRay returns a new instance of Ray at the given position vector pos at a specified radian angle
func (rayCaster *RayCaster) NewRay(pos pixel.Vec, angle float64) *Ray {
	dir := pixel.Vec.Rotated(pos, angle)

	return &Ray{
		RayCaster: *rayCaster,
		pos:       pos,
		dir:       dir,
	}
}

// testing function
func (ray *Ray) lookAt(x, y float64) {
	ray.dir.X = x - ray.pos.X
	ray.dir.Y = y - ray.pos.Y
	ray.dir = ray.dir.Normal()
}

// cast takes a boundary as parameter and then returns the intersection point of the ray with the boundary.
// Returns nil if the ray does not intersect
func (ray *Ray) cast(boundary *Boundary) *pixel.Vec {
	var point pixel.Vec

	x1 := boundary.from.X
	y1 := boundary.from.Y
	x2 := boundary.to.X
	y2 := boundary.to.Y

	x3 := ray.pos.X
	y3 := ray.pos.Y
	x4 := ray.pos.X + ray.dir.X
	y4 := ray.pos.Y + ray.dir.Y

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denominator == 0 {
		return nil
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / denominator
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / denominator

	if t > 0 && t < 1 && u > 0 {
		point.X = x1 + (t * (x2 - x1))
		point.Y = y1 + (t * (y2 - y1))
		return &point
	} else {
		return nil
	}

}

// show renders the ray
func (ray *Ray) show() {
	ray.RayCaster.imd.Color = colornames.Gold
	ray.RayCaster.imd.EndShape = imdraw.SharpEndShape
	ray.RayCaster.imd.Push(ray.pos)
	ray.RayCaster.imd.Push(pixel.V(ray.dir.X*10, ray.dir.Y*10))
	ray.RayCaster.imd.Line(1)
}
