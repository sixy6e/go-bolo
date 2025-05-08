package bolo

const (
	// The time between base getting new units of armour/shells/mines
	BASE_ADD_TIME uint16 = 1670

	// Defines how much the bases can hold
	BASE_FULL_ARMOUR uint8 = 90
	BASE_FULL_SHELLS uint8 = 90
	BASE_FULL_MINES  uint8 = 90

	BASE_NOT_FOUND            uint8  = 254
	BASE_TICKS_BETWEEN_REFUEL uint16 = 1000 // BaseTimer is int32 ...

	// A base is dead if it has 9 armour
	BASE_DEAD uint8 = 9

	// Base will soak shots until it has this much armour
	BASE_MIN_CAN_HIT uint8 = 4

	// This is the units at which bolo will display
	BASE_DISPLAY_X uint8 = 4

	// Minimum amount of stocks a base can have of each item
	BASE_MIN_SHELLS uint8 = 0
	// Base will keep 10 units (2 armour) for itself to defend against attacks
	BASE_MIN_ARMOUR uint8 = 10
	BASE_MIN_MINES  uint8 = 0

	// The amount given to a tank each time
	BASE_SHELLS_GIVE uint8 = 1
	BASE_ARMOUR_GIVE uint8 = 5
	BASE_MINES_GIVE  uint8 = 1

	BASE_REFUEL_ARMOUR uint8   = 46
	BASE_REFUEL_SHELLS float32 = 7.5
	BASE_REFUEL_MINES  float32 = 7.5

	// Minimum armour required for a capture
	MIN_ARMOUR_CAPTURE uint8 = 9

	// Range something has to be in to be shown on the display
	// 7 map squares or 1792 WORLD units
	BASE_STATUS_RANGE uint16 = 1792

	// This is for the half-tick calculator these numbers are simply unique identifiers
	BASES_HALFTICK_TYPE_SHELL  uint8 = 1
	BASES_HALFTICK_TYPE_MINE   uint8 = 2
	BASES_HALFTICK_TYPE_ARMOUR uint8 = 3

	// Bases brain info
	BASES_BRAIN_FRIENDLY    uint8 = 0
	BASES_BRAIN_NEUTRAL     uint8 = 2
	BASES_BRAIN_HOSTILE     uint8 = 1
	BASES_BRAIN_OBJECT_TYPE uint8 = 3

	// Maximum number of bases
	MAX_BASES     uint8  = 16
	SIZE_OF_BASES uint16 = 260
)

type BaseAlliance uint8

const (
	BaseDead BaseAlliance = iota
	BaseOwnGood
	BaseAllyGood
	BaseNeutral
	BaseEvil
)

var BaseTimer = [MAX_TANKS]int32{}

type Base struct {
	// X co-ordinate on the map
	X uint8
	// Y co-ordinate on the map
	Y uint8
	// should be 255 except in speciality maps
	Owner uint8
	// Initial stocks of base. Maximum value of 90
	Armour uint8
	// Initial stocks of base. Maximum value of 90
	Shells uint8
	// Initial stocks of base. Maximum value of 90
	Mines uint8
	// Refuelling time
	RefuelTime uint8
	// Time between stock updates
	BaseTime int32
	// Is this the first time the tank has stopped on the base
	JustStopped bool
}

type Bases struct {
	Items    [MAX_BASES]Base
	NumBases uint8
}

func NewBases() (bases *Bases) {
	var (
		i uint8
	)
	bases = new(Bases)

	bases.NumBases = uint8(0)

	for i = 0; i < MAX_BASES; i++ {
		bases.Items[i].BaseTime = int32(0)
		bases.Items[i].JustStopped = true
	}

	for i = 0; i < MAX_TANKS; i++ {
		BaseTimer[i] = int32(30_000)
	}

	// if(netGetType() == netSingle) {
	//     baseTimer[0]=BASE_TICKS_BETWEEN_REFUEL; // BaseTimer is int32 ...
	// }

	return bases
}

func (bss *Bases) SetNumBases(num uint8) {
	if num <= MAX_BASES {
		bss.NumBases = num
	}
}

// should this be not a pointer given Bases.Items []Base ??
// or turn to be []*Base ??
// c code doesn't do a new Base, but the java code does
// or an options builder???
func NewBase() (base *Base) { // should this be not a pointer given Bases.Items []Base ??
	return base
}

func (bss *Bases) SetBase(base *Base, base_num uint8) {
	if base_num > 0 && base_num <= bss.NumBases {
		// also include at some point
		// if (item->owner > (MAX_TANKS-1) && item->owner != NEUTRAL) {
		//     item->owner = NEUTRAL;
		// }
		bss.Items[base_num].X = base.X
		bss.Items[base_num].Y = base.Y
		bss.Items[base_num].Owner = base.Owner
		bss.Items[base_num].Armour = base.Armour
		bss.Items[base_num].Shells = base.Shells
		bss.Items[base_num].Mines = base.Mines
		bss.Items[base_num].RefuelTime = uint8(0)
		bss.Items[base_num].BaseTime = base.BaseTime
		bss.Items[base_num].JustStopped = true
	}
}
