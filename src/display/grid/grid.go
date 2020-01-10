package grid

//Grid is the grid type
type Grid struct {
	Grid [][]byte
}

//NewGridBySize returns new grid by x and y size
func NewGridBySize(x int, y int) *Grid {
	g := make([][]byte, x)
	for i := range g {
		g[i] = make([]byte, y)
		for j := range g[i] {
			g[i][j] = ' '
		}
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
