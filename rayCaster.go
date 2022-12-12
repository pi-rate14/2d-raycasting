package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
)

type RayCaster struct {
	imd        *imdraw.IMDraw // imd is the imageDraw object responsible for drawing the particle and rays
	boundImd   *imdraw.IMDraw // boundImd is the imageDraw object responsible for drawing the boundaries
	boundaries []Boundary     // boundaries is the list of all the boundaries in the RayCaster
}

// NewRayCaster returns an instance of the RayCaster
func NewRayCaster() *RayCaster {
	return &RayCaster{
		boundImd: imdraw.New(nil),
		imd:      imdraw.New(nil),
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Simple Ray casting",     // Title of the window
		Bounds: pixel.R(0, 0, 960, 1080), // Dimensions of the window
		VSync:  true,                     // Sets the refresh rate of the window equal to the monitor's refresh rate
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	rayCaster := NewRayCaster()

	rayCaster.addRandomBoundaries(cfg)

	rayCaster.addWallBoundaries(cfg)

	particle := rayCaster.NewParticle(win.Bounds().Center())

	for !win.Closed() {

		win.Clear(colornames.Black)

		// rendering all the boundaries
		for _, boundary := range rayCaster.boundaries {
			boundary.show()
			rayCaster.boundImd.Draw(win)
		}

		// updating particle's position based on the keystroke
		if win.JustPressed(pixelgl.KeyLeft) {
			rayCaster.imd.Clear()
			particle.update(-50, 0)
		}
		if win.JustPressed(pixelgl.KeyRight) {
			rayCaster.imd.Clear()
			particle.update(50, 0)
		}
		if win.JustPressed(pixelgl.KeyUp) {
			rayCaster.imd.Clear()
			particle.update(0, 50)
		}
		if win.JustPressed(pixelgl.KeyDown) {
			rayCaster.imd.Clear()
			particle.update(0, -50)
		}
		particle.show()
		rayCaster.imd.Draw(win)
		particle.look()

		win.Update()
	}
}

// randomVector returns a random vector inside the rectangle object r. r is optimally the window.
func randomVector(r pixel.Rect) pixel.Vec {
	return pixel.Vec{
		X: rand.Float64() * r.W(),
		Y: rand.Float64() * r.H(),
	}
}
