package main

import (
	"fmt"
	"strconv"
)

// TODO : Move this code in util's package

type Position struct {
	X byte
	Y uint8
}

// This function get a string in parameter (ex: "J6") and return a Position struct
func GetPositionFromString(inputPos string) Position {
	var pos Position
	YtoInt, _ := strconv.Atoi(inputPos[1:2])
	pos.Y = uint8(YtoInt)

	mapOfCord := map[string]byte{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10}
	pos.X = mapOfCord[inputPos[:1]]

	return pos
}

func main() {

	// Test of the function getPositionFromString
	result := GetPositionFromString("J6")
	fmt.Println(result)

}
