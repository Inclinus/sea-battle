package boats

import (
	"math/rand"
	"time"

	"sea-battle/internal/utils"
)

/*
	Overview of a 4-size boat:
		Clean     : O O O O
		Hit       : O # O O
		Destroyed : # # # #
*/

/*	
	The position is the coordinates of the boat's stern (the back of the boat).
	Since we can't distinguish the stern & the prow of a boat, the position is
	simply one of the extremity of the boat.
*/
type Boat struct {
	Position []utils.Position
	Direction string
	Size uint8
}

func isBoatOverlapping(boat Boat, boats [5]Boat) bool {
	for _, b := range boats {
		for _, p := range b.Position {
			for _, bp := range boat.Position {
				if (p.X == bp.X && p.Y == bp.Y) {
					return true
				}
			}
		}
	}

	return false
}

/*
	Returns an array of 5 boats with random positions & direction

	It verifies that there is no boat overlapping another one & that the number
	of boats of same size doesn't exceed the limit, which is:
		- 1 boat of size 2;
		- 2 boats of size 3;
		- 2 boats of size 3.
*/
func GenerateRandomBoats() (boats [5]Boat) {
	// Seed for randomness
	rand.Seed(time.Now().UnixMicro())

	var directions = [4]string {
		"T", // Top
		"R", // Right
		"B", // Bottom
		"L", // Left
	}

	boatsAmountLimits := map[uint8]uint8{
		2: 1, // Only 1 boat of size 2 can exist
		3: 2, // Only 2 boats of size 3 can exist
		4: 2, // Only 2 boats of size 4 can exist
	}
	boatsCounters := make(map[uint8]uint8)

	// Generate data for each boat
	for i := 0; i < 5; i++ {
		// Generate boat size by checking if it doesn't exceed the limit
		var size uint8
		for {
			size = uint8(2 + rand.Intn(4))

			// Check if there are not too many boats of same size
			if boatsCounters[size] < boatsAmountLimits[size] {
				boatsCounters[size]++
				break
			}
		}

		// Generate direction
		direction := directions[rand.Intn(4)]

		// Generate position
		var position []utils.Position
		// While loop for checking if boat isn't overlapping another one
		for {
			for i := uint8(0); i < size; i++{
				if (i == 0) {
					// Push the first position
					position = append(position, utils.Position{
						X: byte(rand.Intn(10)),
						Y: uint8(rand.Intn(10)),
					})
				} else {
					// Push next positions depending on the direction & the size
					switch direction {
					case "T":
						position = append(position, utils.Position{
							X: position[i-1].X,
							Y: position[i-1].Y + 1,
						})

					case "R":
						position = append(position, utils.Position{
							X: position[i-1].X + 1,
							Y: position[i-1].Y,
						})

					case "B":
						position = append(position, utils.Position{
							X: position[i-1].X,
							Y: position[i-1].Y - 1,
						})

					case "L":
						position = append(position, utils.Position{
							X: position[i-1].X - 1,
							Y: position[i-1].Y,
						})

					default:
						panic("Invalid direction")
					}
				}
			}

			// Create boat
			boat := Boat{position, direction, size}

			// Append boat to the list if it doesn't overlap another one
			// If not, it will regenerate the boat
			// by going back to the beginning of the loop
			if (!isBoatOverlapping(boat, boats)) {
				boats[i] = boat
				break
			}
		}

	}

	return boats
}
