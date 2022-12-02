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
	Position utils.Position
	Direction string
	Size uint8
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
		// Generate position by checking if it's not overlapping another boat
		var position utils.Position
		var direction string
		for {
			position = utils.Position{
				X: byte(rand.Intn(10)),
				Y: uint8(rand.Intn(10)),
			}
			direction = directions[rand.Intn(4)]

			if true { // TODO: boat overlapping verification
				break
			}
		}

		/*
		positions := utils.Position{X: byte(rand.Intn(10)), Y: uint8(rand.Intn(10))}
		direction := directions[rand.Intn(4)]
		*/

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

		// Create boat
		boat := Boat{position, direction, size}

		// Append boat to the list
		boats[i] = boat
	}

	return boats
}
