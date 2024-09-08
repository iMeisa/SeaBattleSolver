package ships

import (
	"SeaBattleSolver/directions"
	"SeaBattleSolver/tiles"
)

var NoneShip = tiles.Tile{
	Type:      tiles.SHIP,
	Direction: directions.NONE,
}

var AnyShip = tiles.Tile{
	Type:      tiles.SHIP,
	Direction: directions.ANY,
}

var UpShip = tiles.Tile{
	Type:      tiles.SHIP,
	Direction: directions.UP,
}

var DownShip = tiles.Tile{
	Type:      tiles.SHIP,
	Direction: directions.DOWN,
}

var LeftShip = tiles.Tile{
	Type:      tiles.SHIP,
	Direction: directions.LEFT,
}

var RightShip = tiles.Tile{
	Type:      tiles.SHIP,
	Direction: directions.RIGHT,
}
