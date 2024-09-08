package main

import (
	"SeaBattleSolver/battle"
	. "SeaBattleSolver/ships"
	. "SeaBattleSolver/tiles"
)

func main() {
	//fmt.Print("Hello, world!")
	battleInit := battle.Battle{
		RowValues:    []int{2, 2, 2, 2, 1, 1},
		ColumnValues: []int{2, 2, 0, 3, 0, 3},
		Board: [][]Tile{
			{EmptyTile, EmptyTile, EmptyTile, EmptyTile, EmptyTile, WaterTile},
			{WaterTile, EmptyTile, EmptyTile, WaterTile, EmptyTile, EmptyTile},
			{EmptyTile, EmptyTile, EmptyTile, NoneShip, EmptyTile, AnyShip},
			{EmptyTile, EmptyTile, EmptyTile, EmptyTile, EmptyTile, EmptyTile},
			{UpShip, EmptyTile, EmptyTile, EmptyTile, EmptyTile, EmptyTile},
			{EmptyTile, EmptyTile, EmptyTile, EmptyTile, EmptyTile, EmptyTile},
		},
	}

	printBoard(battleInit)
}
