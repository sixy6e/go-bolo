package bolo

// TODO; maybe the base type should be uint16???

const (
	// Number of characters in the BMAPBOLO id tag
	LENGTH_ID uint8 = 8
	// Current version
	CURRENT_MAP_VERSION uint64 = 32 // probably not needed

	// Number of bytes in structures (probably not needed)
	SIZEOFBMAP_PILL_INFO  uint8 = 5
	SIZEOFBAMP_BASE_INFO  uint8 = 6
	SIZEOFBMAP_START_INFO uint8 = 3
	SIZEOFBMAP_RUN_HEADER uint8 = 4

	// Map edges for mines
	MAP_MINE_EDGE_LEFT   uint8 = 20
	MAP_MINE_EDGE_RIGHT  uint8 = 236
	MAP_MINE_EDGE_TOP    uint8 = 20
	MAP_MINE_EDGE_BOTTOM uint8 = 236

	// For identical terrain this is the first map code
	MAP_CODE_IDENTICAL_START uint8 = 8
	// Last length
	MAP_CODE_IDENTICAL_END uint8 = 15
	// We want to skip two items for identical map squares
	MAP_CODE_IDENTICAL_SKIP uint8 = 2
	MAP_CODE_DIFFERENT_END  uint8 = 7

	// maps are 256x256 units square
	MAP_ARRAY_SIZE uint16 = 256
	MAP_ARRAY_LAST uint8  = 255

	// All map files should begin with BMAPBOLO (potentially change)
	MAP_HEADER string = "BMAPBOLO"

	// Must shift 4 bits to get a nibble
	MAP_SHIFT_SIZE uint8 = 4

	// Length of six for map to be the same
	MAP_RUN_SAME uint8 = 6
	MAP_RUN_DIFF uint8 = 8

	// Speeds of tanks over various terrains
	// (maybe rename to MAP_TANKSPEED_DEEPSEA etc)
	MAP_SPEED_TDEEPSEA      uint8 = 3 // Tank in deep sea kills tanks
	MAP_SPEED_TBUILDING     uint8 = 0
	MAP_SPEED_TRIVER        uint8 = 3
	MAP_SPEED_TSWAMP        uint8 = 3
	MAP_SPEED_TCRATER       uint8 = 3
	MAP_SPEED_TROAD         uint8 = 16
	MAP_SPEED_TFOREST       uint8 = 6
	MAP_SPEED_TRUBBLE       uint8 = 3
	MAP_SPEED_TGRASS        uint8 = 12
	MAP_SPEED_THALFBUILDING uint8 = 0
	MAP_SPEED_TBOAT         uint8 = 16 // TODO check; original comment states: "was 3 - check ??
	MAP_SPEED_TREFBASE      uint8 = 16 // refueling base
	MAP_SPEED_TPILLBOX      uint8 = 0

	// Speeds of man over various terrains
	MAP_MANSPEED_TDEEPSEA      uint8 = 0
	MAP_MANSPEED_TBUILDING     uint8 = 0
	MAP_MANSPEED_TRIVER        uint8 = 0
	MAP_MANSPEED_TSWAMP        uint8 = 4
	MAP_MANSPEED_TCRATER       uint8 = 4
	MAP_MANSPEED_TROAD         uint8 = 16
	MAP_MANSPEED_TFOREST       uint8 = 8
	MAP_MANSPEED_TRUBBLE       uint8 = 4
	MAP_MANSPEED_TGRASS        uint8 = 16
	MAP_MANSPEED_THALFBUILDING uint8 = 0
	MAP_MANSPEED_TBOAT         uint8 = 16
	MAP_MANSPEED_TREFBASE      uint8 = 16
	MAP_MANSPEED_TPILLBOX      uint8 = 0

	// Speeds of tanks turn given in bradians - 256 radians in a circle
	MAP_TURN_TDEEPSEA      float32 = 0.5
	MAP_TURN_TBUILDING     float32 = 0
	MAP_TURN_TRIVER        float32 = 0.25
	MAP_TURN_TSWAMP        float32 = 0.25
	MAP_TURN_TCRATER       float32 = 0.25
	MAP_TURN_TROAD         float32 = 1
	MAP_TURN_TFOREST       float32 = 0.5
	MAP_TURN_TRUBBLE       float32 = 0.25
	MAP_TURN_TGRASS        float32 = 1
	MAP_TURN_THALFBUILDING float32 = 0
	MAP_TURN_TBOAT         float32 = 1
	MAP_TURN_TREFBASE      float32 = 1
	MAP_TURN_TPILLBOX      float32 = 0

	// The maximum amount of time to wait for the server to authorise this change
	// MAP_MAX_SERVER_WAIT uint8 = 3  // defined in bolo_map.h
	MAP_MAX_SERVER_WAIT uint16 = 200 // redefined in bolo_map.c
)

type MapRunState struct {
	HighLength uint8
	LowLength  uint8
	HighDiff   uint8
	LowDIff    uint8
	HighSame   uint8
	LowSame    uint8
}

type BmapRunHeader struct {
	DataLength uint8
	Y          uint8
	StartX     uint8
	EndX       uint8
}

type BmapRun struct {
	DataLength uint8
	Y          uint8
	StartX     uint8
	EndX       uint8
	Data       [255]uint8
}

// There has to be a better way than a self-referential struct
// even if it is a pointer to another instance of itself
type MapNet struct {
	// Next item
	Next *MapNet
	// Previous item
	Previous *MapNet
	// X Map position
	MapX uint8
	// Y Map position
	MapY uint8
	// New terrain at that position
	Terrain uint8 // we could have a Terrain type, so probs need a different name
	// The old terrain at that position
	OldTerrain uint8
	// Do we need to send this
	NeedSend bool
	// How long have we been waiting for the server?
	Length uint8
}

type Map struct {
	MapItem [][]uint8
	Mn      MapNet
	MnInc   MapNet
}

func NewMap(mp *Map) {
	var (
		y uint16
		x uint16
	)
	mp = new(Map)

	for y = 0; y < MAP_ARRAY_SIZE; y++ {
		for x = 0; x < MAP_ARRAY_SIZE; x++ {
			mp.MapItem[y][x] = uint8(255) // TODO; populate with DEEP_SEA value
		}
	}
}
