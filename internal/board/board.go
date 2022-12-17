package board

import (
	"fmt"

	"sea-battle/internal/boats"
	"sea-battle/internal/shots"
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

	for i := 1; i <= 10; i++ {
		fmt.Println("   -----------------------------------------")
		for j := 0; j <= 10; j++ {
			if j == 0 {
				fmt.Printf("%02d |", i)
			} else {
				symbol := " "

				// Check if there is a boat at this position
				for _, boat := range boats {
					if boat.Position.X == byte(j) && boat.Position.Y == uint8(i) {
						symbol = "O"
					}
				}

				// Check if there is a shot at this position
				for _, shot := range shots {
					if shot.Position.X == byte(j) && shot.Position.Y == uint8(i) {
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
