package display

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/go-grid-vis/src/display/grid"
)

//Display represents a display
type Display struct {
	grid *grid.Grid
	ch   chan int
	sync.Mutex
}

//NewDisplayBySize creates a new rectangular display
func NewDisplayBySize(x, y int) *Display {
	return &Display{
		grid: grid.NewGridBySize(x, y),
	}
}

//NewSquareDisplay creates a new square display of the specified width
func NewSquareDisplay(x int) *Display {
	return &Display{
		grid: grid.NewGridSquare(x),
	}
}

//NewDisplayInitialized creates a grid display from an existing 2d char array
func NewDisplayInitialized(b [][]byte) *Display {
	return &Display{
		grid: grid.NewGridWithInitializer(b),
	}
}

//Start shows the display and updates it automatically
func (d *Display) Start() {
	d.ch = make(chan int, 1)
	d.ch <- 1
	for range d.ch {
		d.print()
	}
}

//Stop stops updating the display
func (d *Display) Stop() {
	close(d.ch)
}

//UpdateAll updates the entire 2d array of characters
func (d *Display) UpdateAll(b [][]byte) {
	d.Lock()
	d.grid.Grid = b
	d.Unlock()
	if d.ch != nil {
		d.ch <- 1
	}
}

//UpdateCoord updates the character at a specific coordinate
func (d *Display) UpdateCoord(x, y int, c byte) {
	d.Lock()
	d.grid.Grid[x][y] = c
	d.Unlock()
	if d.ch != nil {
		d.ch <- 1
	}
}

//Print prints the grid
func (d *Display) print() {
	d.Lock()
	clr := exec.Command("clear")
	clr.Stdout = os.Stdout
	clr.Run()
	for _, row := range d.grid.Grid {
		fmt.Printf("%s\n", row)
	}
	d.Unlock()
}
