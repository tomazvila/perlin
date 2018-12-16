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
	thread_n       = 4
	size           = 1000
)

func calculate_column(from, to int, dc *gg.Context, p *perlin.Perlin, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	for x := 0; x < size; x++ {
		for y := from; y <= to; y++ {
			gradient := math.Abs(p.Noise2D(float64(x)/10, float64(y)/10))
			//fmt.Printf("%0.0d\t%0.0d\t%0.4f\n", x, y, gradient)
			m.Lock()
			dc.SetRGB(gradient, gradient, gradient)
			dc.DrawPoint(float64(x), float64(y), 1.)
			dc.Fill()
			m.Unlock()
		}
	}
}

func main() {
	start := time.Now()
	var waitGroup sync.WaitGroup
	dc := gg.NewContext(size, size)
	p := perlin.NewPerlin(alpha, beta, n, seed)
	var m sync.Mutex
	idk := (int(size) / int(thread_n))
	fmt.Println("idf", idk)
	var x_from int = 0
	var x_to int = idk
	for i := 0; i < thread_n; i++ {
		waitGroup.Add(1)
		fmt.Println("from", x_from)
		fmt.Println("to", x_to-1)
		go calculate_column(x_from, x_to-1, dc, p, &waitGroup, &m)
		x_from = x_to
		x_to += idk
	}
	waitGroup.Wait()
	dc.SavePNG("out.png")

	elapsed := time.Since(start)

	fmt.Println("time took to draw", elapsed)
}
