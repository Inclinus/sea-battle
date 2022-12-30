package board

import (
	"fmt"

	"sea-battle/internal/boats"
	"sea-battle/internal/shots"
	"sea-battle/internal/utils"
)

/*
	Overview of an empty sea battle board:

		A   B   C   D   E   F   G   H   I    J
	   -----------------------------------------
	01 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	02 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	03 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	04 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	05 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	06 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	07 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	08 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	09 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
	10 |   |   |   |   |   |   |   |   |   |   |
	   -----------------------------------------
*/

/*
		Prints an empty board for demonstration purposes (eg: tutorial)

	 	IMPORTANT: if user's terminal is less wide than 44 cols, the board will not
		be printed correctly
*/
func PrintEmptyBoard() {
	fmt.Println("\n     A   B   C   D   E   F   G   H   I   J")

	for i := 1; i <= 10; i++ {
		fmt.Println("   -----------------------------------------")
		fmt.Printf("%02d |   |   |   |   |   |   |   |   |   |   |\n", i)
	}

	fmt.Printf("   -----------------------------------------\n\n")
}

/*
Prints a board with shots & boats

IMPORTANT: if user's terminal is less wide than 44 cols, the board will not
be printed correctly
*/
func PrintBoard(boats [5]boats.Boat, shots []shots.Shot) {
	fmt.Println("\n     A   B   C   D   E   F   G   H   I   J")

	// Concatenate all boats' positions
	var allBoatsPositions []utils.Position
	for _, boat := range boats {
		allBoatsPositions = append(allBoatsPositions, boat.Position...)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("   -----------------------------------------")
		for j := 0; j <= 10; j++ {
			if j == 0 {
				fmt.Printf("%02d |", i)
			} else {
				/*
					Symbols:
					■ -> boat
					0 -> missed shot
					X -> hit shot
					# -> destroyed boat
				*/

				symbol := " "

				// TO DO: print all symbols

				// Check if there is a boat at this position
				for _, boatPosition := range allBoatsPositions {
					if boatPosition.X == uint8(j) && boatPosition.Y == uint8(i) {
						symbol = "■"
					}
				}

				// Check if there is a shot at this position
				for _, shot := range shots {
					if shot.Position.X == uint8(j) && shot.Position.Y == uint8(i) {
						symbol = "X"
					}
				}

				fmt.Printf(" %s |", symbol)
			}
		}
		fmt.Println()
	}

	fmt.Printf("   -----------------------------------------\n\n")
}
