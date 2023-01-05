package board

import (
	"fmt"
	"strconv"

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

var BoatsBoard [5]boats.Boat

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
func PrintBoard(boats [5]boats.Boat, isEnemyBoard bool) {
	fmt.Println("\n     A   B   C   D   E   F   G   H   I   J")

	allShots := *shots.GetAllShots()
	// Get all alive & destroyed boats positions
	var aliveBoatsPositions []utils.Position
	var destroyedBoatsPositions []utils.Position
	for _, boat := range boats {
		if boat.Destroyed {
			destroyedBoatsPositions = append(destroyedBoatsPositions, boat.Position...)
		} else {
			aliveBoatsPositions = append(aliveBoatsPositions, boat.Position...)
		}
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
					O -> missed shot
					X -> hit shot
					# -> destroyed boat
				*/

				symbol := " "

				// Check if there is a boat alive at this position
				if !isEnemyBoard {
					for _, boatPosition := range aliveBoatsPositions {
						if boatPosition.X == uint8(j) && boatPosition.Y == uint8(i) {
							symbol = "■"
						}
					}
				}

				// Check if there is a destroyed boat at this position
				for _, boatPosition := range destroyedBoatsPositions {
					if boatPosition.X == uint8(j) && boatPosition.Y == uint8(i) {
						symbol = "#"
					}
				}

				// Check if there is a shot at this position
				for _, shot := range allShots {
					if shot.Hit && shot.Position.X == uint8(j) && shot.Position.Y == uint8(i) {
						symbol = "X"
					} else if shot.Position.X == uint8(j) && shot.Position.Y == uint8(i) {
						symbol = "O"
					}
				}

				fmt.Printf(" %s |", symbol)
			}
		}
		fmt.Println()
	}

	fmt.Printf("   -----------------------------------------\n\n")
}

// This function get a string in parameter (ex: "J6") and return a Position struct
func GetPositionFromString(inputPos string) utils.Position {
	var pos utils.Position
	YtoInt, _ := strconv.Atoi(inputPos[1:2])
	pos.Y = uint8(YtoInt)

	mapOfCord := map[string]byte{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10}
	pos.X = mapOfCord[inputPos[:1]]

	return pos
}

func InitBoatsBoard(bBoard [5]boats.Boat) {
	BoatsBoard = bBoard
}

func GetBoatsBoard() *[5]boats.Boat {
	return &BoatsBoard
}
