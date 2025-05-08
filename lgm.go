package bolo

const (
	LGM_TREE_REQUEST     uint8 = 0
	LGM_ROAD_REQUEST     uint8 = 1
	LGM_BUILDING_REQUEST uint8 = 2
	LGM_PILL_REQUEST     uint8 = 3
	LGM_MINE_REQUEST     uint8 = 4
	LGM_BOAT_REQUEST     uint8 = 5

	// Man is doing nothing
	LGM_IDLE         uint8 = 6
	LGM_STATE_IDLE   uint8 = 0
	LGM_STATE_GOING  uint8 = 1
	LGM_STATE_RETURN uint8 = 2

	LGM_COST_ROAD           uint8 = 2
	LGM_COST_BUILDING       uint8 = 2
	LGM_COST_REPAIRBUILDING uint8 = 1
	LGM_COST_BOAT           uint8 = 20
	LGM_COST_PILLNEW        uint8 = 4
	LGM_COST_PILLREPAIR     uint8 = 1
	LGM_COST_MINE           uint8 = 1

	LGM_NO_PILL uint8 = 37

	LGM_GATHER_TREE uint8 = 4

	LGM_SIZE_X uint8 = 3
	LGM_SIZE_Y uint8 = 4

	// There are 3 frames in animation 0-3
	LGM_MAX_FRAMES uint8 = 2

	// Blessed area surrounding the tank
	LGM_TANKBOAT_LEAVE  uint8 = 144
	LGM_TANKBOAT_RETURN uint8 = 160

	// Speed the helicpter flys at
	LGM_HELICOPTER_SPEED uint8 = 3
	// The frame number for the helicopter
	LGM_HELICOPTER_FRAME uint8 = 3

	// 20 ticks to builds something
	LGM_BUILD_TIME uint8 = 20

	// Min and Max distance away from thing to achieve goal
	LGM_MIN_GOAL int8  = -16
	LGM_MAX_GOAL uint8 = 16

	// Min and max distance away from return to the tank goal
	// this was set to 16 * 5 which is 80 world coordinates.
	// However, the LGM should be able to return to the tank, from the tanks permiter,
	// so it was changed to 16 * 8.
	LGM_RETURN_MIN_GOAL int16 = -128
	LGM_RETURN_MAX_GOAL int16 = 128

	// The LGMs brain state
	LGM_BRAIN_INTANK uint8 = 0
	LGM_BRAIN_DEAD   uint8 = 1
	// Otherwise
	LGM_BRAIN_MOVING uint8 = 2

	// LGM obstructed state
	LGM_BRAIN_FREE    uint8 = 0
	LGM_BRAIN_PARTIAL uint8 = 1
	LGM_BRAIN_TOTAL   uint8 = 2
)

// LittleGreenMan is the tanks builder.
type LittleGreenMan struct {
	// X World coordinate
	X WORLD
	// Y World coordinate
	Y WORLD
	// X destination (where the man is heading)
	DestX WORLD
	// Y destination
	DestY WORLD
	// Present state (0=idle, 1=going to do work, 2=returning to tank)
	State uint8
	// Is it counting down waiting for something?
	WaitTime uint8
	// Is the building in the tank?
	InTank bool
	// Is the builder dead?
	IsDead bool
	// Number of trees being carried
	NumTrees uint8
	// Number of mines being carried
	NumMines uint8
	// Number of pillboxes being carried
	NumPills uint8
	// What its current task is
	Action uint8
	// X map co-ordinates of next action
	NextX uint8
	// Y map co-ordinates of next action
	NextY uint8
	// Next action to do
	NextAction uint8
	// Animate frame
	Frame uint8
	// The LGM blessed X map square (can travel over it)
	BlessX bool
	// The LGM blessed Y map square (can travel over it)
	BlessY bool
	// Did the LGM land on top of a building/base/pillbox from a parachute?
	OnTop bool
	// For brain - 0=free, 1=touching wall, 2=completely stuck
	Obstructed uint8
	// Player number
	PlayerNum uint8
}

func NewLittleGreenMan(player_num uint8) (lgm *LittleGreenMan) {
	lgm = new(LittleGreenMan)

	lgm.PlayerNum = player_num
	lgm.X = WORLD(0)
	lgm.Y = WORLD(0)
	lgm.InTank = true
	lgm.IsDead = false
	lgm.NextAction = LGM_IDLE
	lgm.Action = LGM_IDLE
	lgm.State = LGM_STATE_IDLE
	lgm.Frame = uint8(0)
	lgm.NumTrees = uint8(0)
	lgm.NumMines = uint8(0)
	lgm.NumPills = LGM_NO_PILL
	lgm.WaitTime = uint8(0)
	lgm.BlessX = false
	lgm.BlessY = false
	lgm.OnTop = false
	lgm.Obstructed = uint8(0)

	return lgm
}
