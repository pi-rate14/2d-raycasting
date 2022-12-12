package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Boundary struct {
	*RayCaster
	from pixel.Vec
	to   pixel.Vec
}

// NewBoundary returns a new boundary between position vectors from and to
func (rayCaster *RayCaster) NewBoundary(from, to pixel.Vec) *Boundary {
	return &Boundary{
		RayCaster: rayCaster,
		from:      from,
		to:        to,
	}
}

// show renders the boundary
func (boundary *Boundary) show() {
	boundary.RayCaster.boundImd.Color = pixel.RGB(1, 1, 1)
	boundary.RayCaster.boundImd.EndShape = imdraw.SharpEndShape
	boundary.RayCaster.boundImd.Push(boundary.from)
	boundary.RayCaster.boundImd.Push(boundary.to)
	boundary.RayCaster.boundImd.Line(5)
}

// addRandomBoundaries adds 5 random boundaries on the window
func (rayCaster *RayCaster) addRandomBoundaries(cfg pixelgl.WindowConfig) {
	for i := 0; i < 5; i++ {
		boundary := rayCaster.NewBoundary(randomVector(cfg.Bounds), randomVector(cfg.Bounds))
		rayCaster.boundaries = append(rayCaster.boundaries, *boundary)
	}
}

// addWallBoundaries adds boundaries on the edges of the window
func (rayCaster *RayCaster) addWallBoundaries(cfg pixelgl.WindowConfig) {
	rayCaster.boundaries = append(rayCaster.boundaries,
		*rayCaster.NewBoundary(
			pixel.V(5, 5),
			pixel.V(cfg.Bounds.W(), 5),
		),
	)

	rayCaster.boundaries = append(rayCaster.boundaries,
		*rayCaster.NewBoundary(
			pixel.V(cfg.Bounds.W(), 0),
			pixel.V(cfg.Bounds.W(), cfg.Bounds.H()),
		),
	)

	rayCaster.boundaries = append(rayCaster.boundaries,
		*rayCaster.NewBoundary(
			pixel.V(cfg.Bounds.W(), cfg.Bounds.H()),
			pixel.V(0, cfg.Bounds.H()),
		),
	)

	rayCaster.boundaries = append(rayCaster.boundaries,
		*rayCaster.NewBoundary(
			pixel.V(0, cfg.Bounds.H()),
			pixel.V(5, 5),
		),
	)
}
