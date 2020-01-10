package main

import (
	"math/rand"
	"time"

	"github.com/go-grid-vis/src/display"
)

func main() {
	grid := display.NewSquareDisplay(50)
	go grid.Start()
	//go func() {
	defer grid.Stop()
	for i := 0; i < 1000; i++ {
		x := rand.Int() % 50
		y := rand.Int() % 50
		grid.UpdateCoord(x, y, (byte(rand.Int()%94) + ' '))
		time.Sleep(17 * time.Millisecond)
	}
	//}()

}
