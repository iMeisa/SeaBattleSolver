package tiles

import "SeaBattleSolver/directions"

type Tile struct {
	Type      TileType
	Direction directions.Direction
}

type TileType int

const (
	NONE TileType = iota
	WATER
	SHIP
)

var EmptyTile = Tile{
	Type: NONE,
}

var WaterTile = Tile{
	Type: WATER,
}
