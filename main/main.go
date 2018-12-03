package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/tomazvila/perlin/perlin"
	"math"
)

const (
	alpha       = 3.
	beta        = 1.5
	n           = 3
	seed  int64 = 100
)

//veliau padaryk, kad nurodzius giju kieki
//sukurtu stulpeliu generavimui po threada

func main() {
	dc := gg.NewContext(1000, 1000)
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 1000; x++ {
		for y := 0.; y < 1000; y++ {
			gradient := math.Abs(p.Noise2D(x/10, y/10))
			fmt.Printf("%0.0f\t%0.0f\t%0.4f\n", x, y, gradient)
			dc.SetRGB(gradient, gradient, gradient)
			dc.DrawPoint(x, y, 1.)
			dc.Fill()
		}
	}
	dc.SavePNG("out.png")
}
