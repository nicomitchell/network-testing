package main

import (
	"math/rand"
	"time"

	"github.com/go-grid-vis/src/display"
)

func main() {
	grid := display.NewSquareDisplay(20)
	go func() {
		defer grid.Stop()
		for i := 0; i < 1000; i++ {
			x := rand.Int() % 20
			y := rand.Int() % 20
			grid.UpdateCoord(x, y, (byte(rand.Int()%94) + ' '))
			time.Sleep(25 * time.Millisecond)
		}
	}()
	grid.Start()

	time.Sleep(1 * time.Second)
}
