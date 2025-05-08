package bolo

// Defines the tile types

const (
	// Horizontal road
	ROAD_HORZ uint8 = 17
	// Vertical road
	ROAD_VERT uint8 = 18

	// The road corners
	ROAD_CORNER1 uint8 = 19
	ROAD_CORNER2 uint8 = 20
	ROAD_CORNER3 uint8 = 21
	ROAD_CORNER4 uint8 = 22
	// The 4 solid road corners
	ROAD_CORNER5 uint8 = 23
	ROAD_CORNER6 uint8 = 24
	ROAD_CORNER7 uint8 = 25
	ROAD_CORNER8 uint8 = 26

	// Solid road
	ROAD_SOLID uint8 = 27
	ROAD_CROSS uint8 = 28

	// T Junctions
	ROAD_T1 uint8 = 29
	ROAD_T2 uint8 = 30
	ROAD_T3 uint8 = 31
	ROAD_T4 uint8 = 32

	// 1 side over other normal
	ROAD_WATER1 uint8 = 33
	ROAD_WATER2 uint8 = 34
	ROAD_WATER3 uint8 = 35
	ROAD_WATER4 uint8 = 36
	// Corners on water
	ROAD_WATER5 uint8 = 37
	ROAD_WATER6 uint8 = 38
	ROAD_WATER7 uint8 = 39
	ROAD_WATER8 uint8 = 40
	// Horizontal water peice
	ROAD_WATER9 uint8 = 41
	// Vertical water peice
	ROAD_WATER10 uint8 = 42
	// Lone water peice
	ROAD_WATER11 uint8 = 43

	// Single side road pieces
	ROAD_SIDE1 uint8 = 44
	ROAD_SIDE2 uint8 = 45
	ROAD_SIDE3 uint8 = 46
	ROAD_SIDE4 uint8 = 47

	// Base states
	// Base in good state
	BASE_GOOD uint8 = 49
	// Base in neutral state
	BASE_NEUTRAL uint8 = 122
	// Base in evil state
	BASE_EVIL uint8 = 163

	// Building tiles
	BUILD_SINGLE uint8 = 50
	BUILD_HORZ   uint8 = 51
	BUILD_VERT   uint8 = 52
	// Building end pieces
	BUILD_HORZEND1 uint8 = 53
	BUILD_HORZEND2 uint8 = 54
	BUILD_VERTEND1 uint8 = 55
	BUILD_VERTEND2 uint8 = 56
	// Building pieces on all sides
	BUILD_SOLID uint8 = 57
	// Building corner pieces (4 corners)
	BUILD_CORNER1 uint8 = 58
	BUILD_CORNER2 uint8 = 59
	BUILD_CORNER3 uint8 = 60
	BUILD_CORNER4 uint8 = 61
	// L shaped pieces
	BUILD_L1 uint8 = 62
	BUILD_L2 uint8 = 63
	BUILD_L3 uint8 = 64
	BUILD_L4 uint8 = 65
	// T shaped pieces
	BUILD_T1 uint8 = 66
	BUILD_T2 uint8 = 67
	BUILD_T3 uint8 = 68
	BUILD_T4 uint8 = 69
	// Cross piece
	BUILD_CROSS uint8 = 70
	// Building side pieces
	BUILD_SIDE1 uint8 = 71
	BUILD_SIDE2 uint8 = 72
	BUILD_SIDE3 uint8 = 73
	BUILD_SIDE4 uint8 = 74
	// Sides and corners
	BUILD_SIDECORN1  uint8 = 75
	BUILD_SIDECORN2  uint8 = 76
	BUILD_SIDECORN3  uint8 = 77
	BUILD_SIDECORN4  uint8 = 78
	BUILD_SIDECORN5  uint8 = 79
	BUILD_SIDECORN6  uint8 = 80
	BUILD_SIDECORN7  uint8 = 81
	BUILD_SIDECORN8  uint8 = 82
	BUILD_SIDECORN9  uint8 = 83
	BUILD_SIDECORN10 uint8 = 84
	BUILD_SIDECORN11 uint8 = 85
	BUILD_SIDECORN12 uint8 = 86
	BUILD_SIDECORN13 uint8 = 87
	BUILD_SIDECORN14 uint8 = 88
	BUILD_SIDECORN15 uint8 = 89
	BUILD_SIDECORN16 uint8 = 90
	// Strange half peices
	BUILD_TWIST1 uint8 = 91
	BUILD_TWIST2 uint8 = 92
	// Most covered
	BUILD_MOST1 uint8 = 93
	BUILD_MOST2 uint8 = 94
	BUILD_MOST3 uint8 = 95
	BUILD_MOST4 uint8 = 96

	// River pieces
	// The four end U pieces
	RIVER_END1 uint8 = 97
	RIVER_END2 uint8 = 98
	RIVER_END3 uint8 = 99
	RIVER_END4 uint8 = 100
	// Solid river piece
	RIVER_SOLID uint8 = 101
	// River surrounded on all sides by buildings
	RIVER_SURROUND uint8 = 102
	// River side pieces eg - | |
	RIVER_SIDE1 uint8 = 103
	RIVER_SIDE2 uint8 = 104
	// One side land others water
	RIVER_ONESIDE1 uint8 = 105
	RIVER_ONESIDE2 uint8 = 106
	RIVER_ONESIDE3 uint8 = 107
	RIVER_ONESIDE4 uint8 = 108
	// Four corner L shaped pieces
	RIVER_CORN1 uint8 = 109
	RIVER_CORN2 uint8 = 110
	RIVER_CORN3 uint8 = 111
	RIVER_CORN4 uint8 = 112

	// Deep Sea Peices
	DEEP_SEA_SOLID uint8 = 113
	// Corner pieces L shaped
	DEEP_SEA_CORN1 uint8 = 114
	DEEP_SEA_CORN2 uint8 = 115
	DEEP_SEA_CORN3 uint8 = 116
	DEEP_SEA_CORN4 uint8 = 117
	// Four sides
	DEEP_SEA_SIDE1 uint8 = 118
	DEEP_SEA_SIDE2 uint8 = 119
	DEEP_SEA_SIDE3 uint8 = 120
	DEEP_SEA_SIDE4 uint8 = 121

	// Tank positions
	TANK_SELF_0  uint8 = 0
	TANK_SELF_1  uint8 = 1
	TANK_SELF_2  uint8 = 2
	TANK_SELF_3  uint8 = 3
	TANK_SELF_4  uint8 = 4
	TANK_SELF_5  uint8 = 5
	TANK_SELF_6  uint8 = 6
	TANK_SELF_7  uint8 = 7
	TANK_SELF_8  uint8 = 8
	TANK_SELF_9  uint8 = 9
	TANK_SELF_10 uint8 = 10
	TANK_SELF_11 uint8 = 11
	TANK_SELF_12 uint8 = 12
	TANK_SELF_13 uint8 = 13
	TANK_SELF_14 uint8 = 14
	TANK_SELF_15 uint8 = 15
	// Tank boad positions
	TANK_SELFBOAT_0  uint8 = 16
	TANK_SELFBOAT_1  uint8 = 17
	TANK_SELFBOAT_2  uint8 = 18
	TANK_SELFBOAT_3  uint8 = 19
	TANK_SELFBOAT_4  uint8 = 20
	TANK_SELFBOAT_5  uint8 = 21
	TANK_SELFBOAT_6  uint8 = 22
	TANK_SELFBOAT_7  uint8 = 23
	TANK_SELFBOAT_8  uint8 = 24
	TANK_SELFBOAT_9  uint8 = 25
	TANK_SELFBOAT_10 uint8 = 26
	TANK_SELFBOAT_11 uint8 = 27
	TANK_SELFBOAT_12 uint8 = 28
	TANK_SELFBOAT_13 uint8 = 29
	TANK_SELFBOAT_14 uint8 = 30
	TANK_SELFBOAT_15 uint8 = 31
	// Tank good states
	TANK_GOOD_0  uint8 = 32
	TANK_GOOD_1  uint8 = 33
	TANK_GOOD_2  uint8 = 34
	TANK_GOOD_3  uint8 = 35
	TANK_GOOD_4  uint8 = 36
	TANK_GOOD_5  uint8 = 37
	TANK_GOOD_6  uint8 = 38
	TANK_GOOD_7  uint8 = 39
	TANK_GOOD_8  uint8 = 40
	TANK_GOOD_9  uint8 = 41
	TANK_GOOD_10 uint8 = 42
	TANK_GOOD_11 uint8 = 43
	TANK_GOOD_12 uint8 = 44
	TANK_GOOD_13 uint8 = 45
	TANK_GOOD_14 uint8 = 46
	TANK_GOOD_15 uint8 = 47
	// Tank good boat states
	TANK_GOODBOAT_0  uint8 = 48
	TANK_GOODBOAT_1  uint8 = 49
	TANK_GOODBOAT_2  uint8 = 50
	TANK_GOODBOAT_3  uint8 = 51
	TANK_GOODBOAT_4  uint8 = 52
	TANK_GOODBOAT_5  uint8 = 53
	TANK_GOODBOAT_6  uint8 = 54
	TANK_GOODBOAT_7  uint8 = 55
	TANK_GOODBOAT_8  uint8 = 56
	TANK_GOODBOAT_9  uint8 = 57
	TANK_GOODBOAT_10 uint8 = 58
	TANK_GOODBOAT_11 uint8 = 59
	TANK_GOODBOAT_12 uint8 = 60
	TANK_GOODBOAT_13 uint8 = 61
	TANK_GOODBOAT_14 uint8 = 62
	TANK_GOODBOAT_15 uint8 = 63
	// Tank evil states
	TANK_EVIL_0  uint8 = 64
	TANK_EVIL_1  uint8 = 65
	TANK_EVIL_2  uint8 = 66
	TANK_EVIL_3  uint8 = 67
	TANK_EVIL_4  uint8 = 68
	TANK_EVIL_5  uint8 = 69
	TANK_EVIL_6  uint8 = 70
	TANK_EVIL_7  uint8 = 71
	TANK_EVIL_8  uint8 = 72
	TANK_EVIL_9  uint8 = 73
	TANK_EVIL_10 uint8 = 74
	TANK_EVIL_11 uint8 = 75
	TANK_EVIL_12 uint8 = 76
	TANK_EVIL_13 uint8 = 77
	TANK_EVIL_14 uint8 = 78
	TANK_EVIL_15 uint8 = 79
	// Tank evil boat states
	TANK_EVILBOAT_0  uint8 = 80
	TANK_EVILBOAT_1  uint8 = 81
	TANK_EVILBOAT_2  uint8 = 82
	TANK_EVILBOAT_3  uint8 = 83
	TANK_EVILBOAT_4  uint8 = 84
	TANK_EVILBOAT_5  uint8 = 85
	TANK_EVILBOAT_6  uint8 = 86
	TANK_EVILBOAT_7  uint8 = 87
	TANK_EVILBOAT_8  uint8 = 88
	TANK_EVILBOAT_9  uint8 = 89
	TANK_EVILBOAT_10 uint8 = 90
	TANK_EVILBOAT_11 uint8 = 91
	TANK_EVILBOAT_12 uint8 = 92
	TANK_EVILBOAT_13 uint8 = 93
	TANK_EVILBOAT_14 uint8 = 94
	TANK_EVILBOAT_15 uint8 = 95
	// Add 16 to get to the begining of the allied tank tiles
	TANK_BOAT_ADD uint8 = 16
	// Add 32 to get to the begining of the allied tank tiles
	TANK_GOOD_ADD uint8 = 32
	// Add 64 to get to the begining of the allied tank tiles
	TANK_EVIL_ADD uint8 = 64
	// Tank transparent
	TANK_TRANSPARENT uint8 = 255

	// Shells and explosions
	SHELL_EXPLOSION8 uint8 = 1
	SHELL_EXPLOSION7 uint8 = 2
	SHELL_EXPLOSION6 uint8 = 3
	SHELL_EXPLOSION5 uint8 = 4
	SHELL_EXPLOSION4 uint8 = 5
	SHELL_EXPLOSION3 uint8 = 6
	SHELL_EXPLOSION2 uint8 = 7
	SHELL_EXPLOSION1 uint8 = 8
	SHELL_DIR0       uint8 = 9
	SHELL_DIR1       uint8 = 10
	SHELL_DIR2       uint8 = 11
	SHELL_DIR3       uint8 = 12
	SHELL_DIR4       uint8 = 13
	SHELL_DIR5       uint8 = 14
	SHELL_DIR6       uint8 = 15
	SHELL_DIR7       uint8 = 16
	SHELL_DIR8       uint8 = 17
	SHELL_DIR9       uint8 = 18
	SHELL_DIR10      uint8 = 19
	SHELL_DIR11      uint8 = 20
	SHELL_DIR12      uint8 = 21
	SHELL_DIR13      uint8 = 22
	SHELL_DIR14      uint8 = 23
	SHELL_DIR15      uint8 = 24

	// Evil pillbox numbers
	PILL_EVIL_15 uint8 = 48
	PILL_EVIL_14 uint8 = 123
	PILL_EVIL_13 uint8 = 124
	PILL_EVIL_12 uint8 = 125
	PILL_EVIL_11 uint8 = 126
	PILL_EVIL_10 uint8 = 127
	PILL_EVIL_9  uint8 = 128
	PILL_EVIL_8  uint8 = 129
	PILL_EVIL_7  uint8 = 130
	PILL_EVIL_6  uint8 = 131
	PILL_EVIL_5  uint8 = 132
	PILL_EVIL_4  uint8 = 133
	PILL_EVIL_3  uint8 = 134
	PILL_EVIL_2  uint8 = 135
	PILL_EVIL_1  uint8 = 136
	PILL_EVIL_0  uint8 = 137
	// Good pillbox numbers
	PILL_GOOD_15 uint8 = 147
	PILL_GOOD_14 uint8 = 148
	PILL_GOOD_13 uint8 = 149
	PILL_GOOD_12 uint8 = 150
	PILL_GOOD_11 uint8 = 151
	PILL_GOOD_10 uint8 = 152
	PILL_GOOD_9  uint8 = 153
	PILL_GOOD_8  uint8 = 154
	PILL_GOOD_7  uint8 = 155
	PILL_GOOD_6  uint8 = 156
	PILL_GOOD_5  uint8 = 157
	PILL_GOOD_4  uint8 = 158
	PILL_GOOD_3  uint8 = 159
	PILL_GOOD_2  uint8 = 160
	PILL_GOOD_1  uint8 = 161
	PILL_GOOD_0  uint8 = 162

	// Forests
	FOREST_SINGLE uint8 = 164
	FOREST_BR     uint8 = 165
	FOREST_BL     uint8 = 166
	FOREST_AL     uint8 = 167
	FOREST_AR     uint8 = 168
	FOREST_ABOVE  uint8 = 169
	FOREST_BELOW  uint8 = 170
	FOREST_LEFT   uint8 = 171
	FOREST_RIGHT  uint8 = 172

	// Craters
	CRATER_SINGLE uint8 = 173
	CRATER_BR     uint8 = 174
	CRATER_BL     uint8 = 175
	CRATER_AL     uint8 = 176
	CRATER_AR     uint8 = 177
	CRATER_ABOVE  uint8 = 178
	CRATER_BELOW  uint8 = 179
	CRATER_LEFT   uint8 = 180
	CRATER_RIGHT  uint8 = 181

	// Boats
	BOAT_0 uint8 = 138
	BOAT_1 uint8 = 139
	BOAT_2 uint8 = 140
	BOAT_3 uint8 = 141
	BOAT_4 uint8 = 142
	BOAT_5 uint8 = 143
	BOAT_6 uint8 = 144
	BOAT_7 uint8 = 145
	BOAT_8 uint8 = 146

	// LGM
	LGM0 uint8 = 0
	LGM1 uint8 = 1
	LGM2 uint8 = 2
	LGM3 uint8 = 3
)
