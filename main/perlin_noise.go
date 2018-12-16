package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/tomazvila/perlin/perlin"
	"math"
	"sync"
	"time"
)

const (
	alpha          = 3.
	beta           = 1.5
	n              = 3
	seed     int64 = 100
	thread_n       = 1
	size           = 10000
)

func calculate_column(from, to int, dc *gg.Context, p *perlin.Perlin, wg *sync.WaitGroup, gradient *[size][size]float64) {
	defer wg.Done()
	for x := 0; x < size; x++ {
		for y := from; y <= to; y++ {
			gradient[int(x)][int(y)] = math.Abs(p.Noise2D(float64(x)/10, float64(y)/10))
		}
	}
}

func main() {
	var waitGroup sync.WaitGroup
	dc := gg.NewContext(size, size)
	p := perlin.NewPerlin(alpha, beta, n, seed)
	thread_size := (int(size) / int(thread_n))
	var x_from int = 0
	var x_to int = thread_size
	gradient := [size][size]float64{}
	start := time.Now()
	for i := 0; i < thread_n; i++ {
		waitGroup.Add(1)
		fmt.Println("from", x_from)
		fmt.Println("to", x_to-1)
		go calculate_column(x_from, x_to-1, dc, p, &waitGroup, &gradient)
		x_from = x_to
		x_to += thread_size
	}
	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Println("time took to draw", elapsed)
	for x := 0.; x < size; x++ {
		for y := 0.; y < size; y++ {
			g := gradient[int(x)][int(y)]
			dc.SetRGB(g, g, g)
			dc.DrawPoint(x, y, 1)
			dc.Fill()
		}
	}
	dc.SavePNG("out.png")
}
