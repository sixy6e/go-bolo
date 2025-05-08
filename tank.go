package bolo

const (
	// TODO; change these to int32 as #define TANK_SHIFT_PIXELSIZE 12 would be int
	// according to google searches
	// #define BASE_VALUE 10
	// #define MAX_VALUE 100U
	// #define BIG_VALUE 10000000000L
	// In this case, BASE_VALUE will be of type int, MAX_VALUE of type unsigned int, and BIG_VALUE of type long

	// Shift world co-ordinate 8 bits right to get the map co-ordinate
	TANK_SHIFT_MAPSIZE int32 = 8
	// Shift world co-ordinate 12 bits right to get map co-ordinate
	TANK_SHIFT_PIXELSIZE int32 = 12
	// Shift things right 2 for reading to the world
	TANK_SHIFT_RIGHT2 int32 = 4
	// There are 16 tank frames (or viewing angles the tank can take)
	TANK_FRAMES int32 = 16
	// How many world co-ordinates to move per tank slide update
	TANK_SLIDE int32 = 16
	// The number of times to update the tank slide in world co-ordinates.
	// The actual number is this +1 becuase it activates tank slide once right away
	TANK_SLIDE_TICKS int32 = 7
	// How many world co-ordinates to move per tank bump update
	TANK_BUMP int32 = 7
	// The number of times to update the tank slide, in ticks.
	// The actual number is this +1 becuase it activates tank slide once right away
	TANK_BUMP_TICKS int32 = 7
	TANK_DEATH_WAIT int32 = 255
	// The acceleration rate of the tank
	TANK_ACCELERATE_RATE float32 = 0.25
	// The deceleration rate of the tank
	TANK_TERRAIN_DECEL_RATE float32 = 0.25
	// The slow down button pressed rate
	TANK_SLOWKEY_RATE float32 = 0.25
	// Auto slowdown deceleration
	TANK_AUTOSLOW_SPEED float32 = 0.25

	// BRadians; 256 in a circle
	// Compass north in terms of tank BRadians
	TANK_NORTH int32 = 0
	// Compass east in terms of tank BRadians
	TANK_EAST int32 = 64
	// Compass south in terms of tank BRadians
	TANK_SOUTH int32 = 128
	// Compass west in terms of tank BRadians
	TANK_WEST int32 = 192

	// Amount to subtract
	TANK_SUBTRACT int32 = 128
	// The square test to see if the tank is hit
	TANK_SUBTRACT_SHELL_HIT_X1 int32 = 144
	TANK_SUBTRACT_SHELL_HIT_X2 int32 = 128
	TANK_SUBTRACT_SHELL_HIT_Y  int32 = 128

	// Minimum length for the gunsight (in half map squares)
	GUNSIGHT_MIN int32 = 2
	// Maximum length for the gunsight (in half map squares)
	GUNSIGHT_MAX int32 = 14
	// Game ticks for the tank to reload
	TANK_RELOAD_TIME int32 = 13
	// Time it takes for the tank to lose some stuff in water
	TANK_WATER_TIME int32 = 15
	// Shells + mines myst be greater than 20 for a big explosion
	TANK_BIG_EXPLOSION_THRESHOLD int32 = 20
	// Amount of damage a mine does to the tank
	MINE_DAMAGE int32 = 10
	// Number of map squares around the tank position that a mine can hurt a tank
	MINE_DAMAGE_DISTANCE_LEFT  int32 = 0 // was -1
	MINE_DAMAGE_DISTANCE_RIGHT int32 = 0 // was -1
	// Tank slows down a speed unit if it hits a wall
	TANK_WALL_SLOW_DOWN int32 = 1
	// Amount to move for checking
	TANK_MOVE_BOAT_SUB int32 = 64
	TANK_MOVE_LAND_SUB int32 = 128
	// Speed we exit the boat at
	BOAT_EXIT_SPEED int32 = 16
	// Minimum distance for seeing a tank in trees = 3 map squares or 768 world co-ordinates
	MIN_TREEHIDE_DIST int32 = 768

	// Max tanks in the game
	MAX_TANKS uint8 = 16
)

const (
	// Tank numbers 1-16
	TANK_0 uint8 = iota // original c code defined this, but maybe it is ignored
	TANK_1
	TANK_2
	TANK_3
	TANK_4
	TANK_5
	TANK_6
	TANK_7
	TANK_8
	TANK_9
	TANK_10
	TANK_11
	TANK_12
	TANK_13
	TANK_14
	TANK_15
	TANK_16 // 16

)

type TankHit uint8
type TankButton uint8

const (
	// Tank has been missed
	TH_Missed TankHit = iota
	// Tank has been hit but not killed
	TH_Hit
	// The tank has been hit, killed, and isn't carrying a lot of stuff so a small explosion
	TH_KillSmall
	// The tank has been hit, killed, and is carrying a lot of stuff so a big explosion
	TH_KillBig
)

const (
	// No buttons being pressed
	T_None TankButton = iota
	// Left button is being pressed
	T_Left
	// Right button is being pressed
	T_Right
	// Accelerate button is being pressed
	T_Accelerate
	// Decelerate button is being pressed
	T_Decelerate
	// Left + Accelerate
	T_LeftAccelerate
	// Right + Accelerate
	T_RightAccelerate
	// Left + Decelerate
	T_LeftDecelerate
	// Right + Decelerate
	T_RIghtDecelerate
)

type VectorBody struct {
	Mass           uint8
	Angle          float32
	AnglePrevious  float32
	SpeedX         float32
	SpeedY         float32
	SpeedXPrevious float32
	SpeedYPrevious float32
}

type TODO uint8 // For unknowns; populate fields now, fix later

// The world co-ordinate for tank - It is a 16 bit number
// Top 8 bits (or byte) represent position on the map
// Next 2 bits represent the tanks sub-map position in pixels (0-15)
// Last 2 bits sub 40 pixels
type WORLD uint16

type Tank struct {
	// X World co-ordinate
	X WORLD
	// Y World co-ordinate
	Y WORLD
	// Amount of armour in tank. Maximum value of 45
	Armour uint8
	// Amount of shells in tank. Maximum value of 45
	Shells uint8
	// Amount of mines in tank. Maximum value of 45
	Mines uint8
	// Tank speed
	Speed float32
	// Amount of trees in tank
	Trees uint8
	// The anlge the tank is pointing on 0-256
	Angle float32
	// Reload tick 0 is OK to shoot else counts back
	Reload uint8
	// Is the tank on a boat?
	OnBoat bool
	// Is the gunsight on or not?
	ShowSight bool
	// Length of the gunsight measure in map units
	SightLength uint8
	// Number of kills the tank has had
	NumKills int32
	// Number of deaths the tank has had
	NumDeaths int32
	// How long it is going to be on the screen till it refreshes
	DeathWait uint8
	// Count for bubbles
	WaterCount uint8
	// Used by brains. Did the tank hit anything
	Obstructed bool
	// Is this a new tank or not (i.e. just died)
	NewTank bool
	// Do we use autoslowdown or not?
	AutoSlowdown bool
	// Auto show/hide of gunsight enabled/disabled
	AutoHideGunsight bool
	// Did the tank just fire?
	JustFired bool
	// Number of times a tank has been hit to determine if they are cheating
	TankHitCount uint8
	// CRC used to detect memory cheats
	Crc int32
	// World X co-ordinate at last tick
	XPrevious WORLD
	// World Y co-ordinate at last tick
	YPrevious WORLD
	// World X co-ordinate at tick before last tick
	XPreviousPrevious WORLD
	// World Y co-ordinate at tick before last tick
	YPreviousPrevious WORLD
	// The pillboxes being carried
	CarryPillboxes TODO
	// This will be used to determine how long the tank should slide when hit
	TankSlideTimer uint8
	// This holds the angle at which the tank was hit with a shell
	TankSlideAngle float32
	// How did the most recent death to the tank occur?
	LastTankDeath uint8
	// Holds tank's actual moving direction and component vectors (x and y axis speed)
	VectorBodyTank VectorBody
	// Holds physics stuff for what hit the tank
	VectorBodyCollide VectorBody
}

func NewTank() (tnk *Tank) {
	tnk = new(Tank)

	tnk.X = WORLD(0) // TODO; check type, init appears to be BYTE
	tnk.Y = WORLD(0)

	return tnk
}

func (tnk *Tank) InWater() {
	var modifications bool = false

	if tnk.Shells > 0 {
		tnk.Shells -= 1
		// tankRegisterChangeByte(value, CRC_SHELLS_OFFSET, (*value)->shells);
		modifications = true
	}

	if tnk.Mines > 0 {
		tnk.Mines -= 1
		// tankRegisterChangeByte(value, CRC_MINES_OFFSET, (*value)->mines);
		modifications = true
	}

	if modifications {
		// update status bars and play bubbles sound
		// frontEndUpdateTankStatusBars((*value)->shells, (*value)->mines, (*value)->armour, (*value)->       trees);
		// frontEndPlaySound(bubbles);
	}
}

// IsMoving confirms whether or not the tank is in motion or not
func (tnk *Tank) IsMoving() (moving bool) {
	if tnk.Speed != 0 {
		moving = true
	}

	return moving
}

func (tnk *Tank) Direction() (direction uint8) {
	// need a func to convert BRadians to 16 direction for both shells and tanks
	// utilGetDir
	return direction
}

func (tnk *Tank) TravelAngle() (tangle uint8) {
	// utilGet16Dir
	return tangle
}

// Direction256 returns the tank direction 0-255
func (tnk *Tank) Direction256() (direction uint8) {
	direction = uint8(tnk.Angle)
	return direction
}
