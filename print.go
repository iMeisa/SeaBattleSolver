package main

import (
	"SeaBattleSolver/battle"
	"SeaBattleSolver/tiles"
	"fmt"
)

func printBoard(battle battle.Battle) {
	fmt.Println()
	fmt.Println()

	for i, row := range battle.Board {

		for _, tile := range row {
			tileDisplay := "-"

			switch tile.Type {
			case tiles.WATER:
				tileDisplay = "W"
			case tiles.SHIP:
				tileDisplay = "S"
			default:
				tileDisplay = "-"
			}

			fmt.Print(tileDisplay + "   ")
		}

		fmt.Println(battle.RowValues[i])
		fmt.Println()
	}

	for _, col := range battle.ColumnValues {
		fmt.Printf("%d   ", col)
	}
}
