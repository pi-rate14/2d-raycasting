package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
