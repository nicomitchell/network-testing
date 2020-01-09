package grid

import "fmt"

//Grid is the grid type
type Grid struct {
	Grid [][]byte
}

//NewGridBySize returns new grid by x and y size
func NewGridBySize(x int, y int) *Grid {
	g := make([][]byte, x)
	for i := range g {
		g[i] = make([]byte, y)
	}
	return &Grid{
		Grid: g,
	}
}

//NewGridWithInitializer uses an existing 2d slice to make a grid
func NewGridWithInitializer(g [][]byte) *Grid {
	return &Grid{Grid: g}
}

//NewGridSquare returns a square grid
func NewGridSquare(x int) *Grid {
	return NewGridBySize(x, x)
}

//Display prints the grid
func (g *Grid) Display() {
	for _, row := range g.Grid {
		fmt.Println(row)
	}
}
