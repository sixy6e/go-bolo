package bolo

const (
	MINES_ARRAY_SIZE uint16 = 256
	MINE_START       uint8  = 10
	MINE_END         uint8  = 15
	MINE_SUBTRACT    uint8  = 8
)

type Mines struct {
	Position    [MINES_ARRAY_SIZE][MINES_ARRAY_SIZE]bool
	HiddenMines bool
}

func NewMines(allow_hidden bool) (mn *Mines) {
	mn = new(Mines)
	mn.HiddenMines = allow_hidden

	return mn
}

func (mn *Mines) Add(x, y uint8) {
	// original c code had [x][y]
	// as the map is square it doesn't matter
	// but if we make non-square, the general convention
	// would be going down the list of slices (y) then across to cell (x)
	// also I'm coming from Python & NumPy land, and c ordering has the fastest moving axis
	// is on the far right (traverse through X before moving to the next row)
	// so i'll stick with the [y, x] logic
	mn.Position[y][x] = true
}

func (mn *Mines) Remove(x, y uint8) {
	mn.Position[y][x] = false
}

func (mn *Mines) Exist(x, y uint8) (exists bool) {
	if x <= MAP_MINE_EDGE_LEFT || x >= MAP_MINE_EDGE_RIGHT || y <= MAP_MINE_EDGE_TOP || y >= MAP_MINE_EDGE_BOTTOM {
		exists = true
	} else if mn.HiddenMines == true {
		exists = mn.Position[y][x]
	} else {
		// exists = screenMapIsMine(xValue, yValue)
	}

	return exists
}
