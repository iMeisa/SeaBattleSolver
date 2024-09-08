package battle

import (
	"SeaBattleSolver/directions"
	"SeaBattleSolver/tiles"
)

type Battle struct {
	RowValues    []int
	ColumnValues []int
	Board        [][]tiles.Tile
}

func (b *Battle) FillWater() {
	var waterCoordinates [][]int
}

func (b *Battle) findNextShip(row, column int) {
	for rowIndex, boardRow := range b.Board {

	}
}

func (b *Battle) crawlShip(row, column int) {
	if b.Board[row][column].Direction == directions.NONE {
		return
	}

}
