package main

import (
	"fmt"
	"strconv"
)

type Position struct {
	X byte
	Y uint8
}

func getPos(inputPos string) Position {
	mapOfCord := map[string]Position{
		"A1":  Position{X: 1, Y: 1},
		"A2":  Position{X: 1, Y: 2},
		"A3":  Position{X: 1, Y: 3},
		"A4":  Position{X: 1, Y: 4},
		"A5":  Position{X: 1, Y: 5},
		"A6":  Position{X: 1, Y: 6},
		"A7":  Position{X: 1, Y: 7},
		"A8":  Position{X: 1, Y: 8},
		"A9":  Position{X: 1, Y: 9},
		"A10": Position{X: 1, Y: 10},

		"B1":  Position{X: 2, Y: 1},
		"B2":  Position{X: 2, Y: 2},
		"B3":  Position{X: 2, Y: 3},
		"B4":  Position{X: 2, Y: 4},
		"B5":  Position{X: 2, Y: 5},
		"B6":  Position{X: 2, Y: 6},
		"B7":  Position{X: 2, Y: 7},
		"B8":  Position{X: 2, Y: 8},
		"B9":  Position{X: 2, Y: 9},
		"B10": Position{X: 2, Y: 10},

		"C1":  Position{X: 3, Y: 1},
		"C2":  Position{X: 3, Y: 2},
		"C3":  Position{X: 3, Y: 3},
		"C4":  Position{X: 3, Y: 4},
		"C5":  Position{X: 3, Y: 5},
		"C6":  Position{X: 3, Y: 6},
		"C7":  Position{X: 3, Y: 7},
		"C8":  Position{X: 3, Y: 8},
		"C9":  Position{X: 3, Y: 9},
		"C10": Position{X: 3, Y: 10},

		"D1":  Position{X: 4, Y: 1},
		"D2":  Position{X: 4, Y: 2},
		"D3":  Position{X: 4, Y: 3},
		"D4":  Position{X: 4, Y: 4},
		"D5":  Position{X: 4, Y: 5},
		"D6":  Position{X: 4, Y: 6},
		"D7":  Position{X: 4, Y: 7},
		"D8":  Position{X: 4, Y: 8},
		"D9":  Position{X: 4, Y: 9},
		"D10": Position{X: 4, Y: 10},

		"E1":  Position{X: 5, Y: 1},
		"E2":  Position{X: 5, Y: 2},
		"E3":  Position{X: 5, Y: 3},
		"E4":  Position{X: 5, Y: 4},
		"E5":  Position{X: 5, Y: 5},
		"E6":  Position{X: 5, Y: 6},
		"E7":  Position{X: 5, Y: 7},
		"E8":  Position{X: 5, Y: 8},
		"E9":  Position{X: 5, Y: 9},
		"E10": Position{X: 5, Y: 10},

		"F1":  Position{X: 6, Y: 1},
		"F2":  Position{X: 6, Y: 2},
		"F3":  Position{X: 6, Y: 3},
		"F4":  Position{X: 6, Y: 4},
		"F5":  Position{X: 6, Y: 5},
		"F6":  Position{X: 6, Y: 6},
		"F7":  Position{X: 6, Y: 7},
		"F8":  Position{X: 6, Y: 8},
		"F9":  Position{X: 6, Y: 9},
		"F10": Position{X: 6, Y: 10},

		"G1":  Position{X: 7, Y: 1},
		"G2":  Position{X: 7, Y: 2},
		"G3":  Position{X: 7, Y: 3},
		"G4":  Position{X: 7, Y: 4},
		"G5":  Position{X: 7, Y: 5},
		"G6":  Position{X: 7, Y: 6},
		"G7":  Position{X: 7, Y: 7},
		"G8":  Position{X: 7, Y: 8},
		"G9":  Position{X: 7, Y: 9},
		"G10": Position{X: 7, Y: 10},

		"H1":  Position{X: 8, Y: 1},
		"H2":  Position{X: 8, Y: 2},
		"H3":  Position{X: 8, Y: 3},
		"H4":  Position{X: 8, Y: 4},
		"H5":  Position{X: 8, Y: 5},
		"H6":  Position{X: 8, Y: 6},
		"H7":  Position{X: 8, Y: 7},
		"H8":  Position{X: 8, Y: 8},
		"H9":  Position{X: 8, Y: 9},
		"H10": Position{X: 8, Y: 10},

		"I1":  Position{X: 9, Y: 1},
		"I2":  Position{X: 9, Y: 2},
		"I3":  Position{X: 9, Y: 3},
		"I4":  Position{X: 9, Y: 4},
		"I5":  Position{X: 9, Y: 5},
		"I6":  Position{X: 9, Y: 6},
		"I7":  Position{X: 9, Y: 7},
		"I8":  Position{X: 9, Y: 8},
		"I9":  Position{X: 9, Y: 9},
		"I10": Position{X: 9, Y: 10},

		"J1":  Position{X: 10, Y: 1},
		"J2":  Position{X: 10, Y: 2},
		"J3":  Position{X: 10, Y: 3},
		"J4":  Position{X: 10, Y: 4},
		"J5":  Position{X: 10, Y: 5},
		"J6":  Position{X: 10, Y: 6},
		"J7":  Position{X: 10, Y: 7},
		"J8":  Position{X: 10, Y: 8},
		"J9":  Position{X: 10, Y: 9},
		"J10": Position{X: 10, Y: 10}}
	return mapOfCord[inputPos]
}

func getPos2(inputPos string) Position {
	var pos Position
	YtoInt, _ := strconv.Atoi(inputPos[1:2])
	pos.Y = uint8(YtoInt)

	mapOfCord := map[string]byte{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10}
	pos.X = mapOfCord[inputPos[:1]]

	return pos
}

func main() {
	result := getPos2("J6")
	fmt.Println(result)
}
