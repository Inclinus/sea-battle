package main

import (
	"fmt"

	"sea-battle/internal/board"
	"sea-battle/internal/boats"
	"sea-battle/internal/ip"
	"sea-battle/internal/shots"
	"sea-battle/internal/utils"
)

func main() {
	ip, port := ip.SplitIpAndPort("192.168.1.1:80")
	fmt.Printf("IP: %s\nPort: %d\n", ip, port)

	boats := boats.GenerateRandomBoats()
	for i, boat := range boats {
		fmt.Printf("Boat %d:\n", i)
		fmt.Printf("\tPosition: %v\n", boat.Position)
		fmt.Printf("\tDirection: %s\n", boat.Direction)
		fmt.Printf("\tSize: %d\n\n", boat.Size)
	}

	// Create an array of allShots
	var allShots []shots.Shot
	allShots = append(allShots, shots.Shot{Position: utils.Position{X: 3, Y: 1}, Hit: true})
	allShots = append(allShots, shots.Shot{Position: utils.Position{X: 9, Y: 2}, Hit: false})
	allShots = append(allShots, shots.Shot{Position: utils.Position{X: 2, Y: 3}, Hit: true})
	allShots = append(allShots, shots.Shot{Position: utils.Position{X: 5, Y: 4}, Hit: false})
	allShots = append(allShots, shots.Shot{Position: utils.Position{X: 10, Y: 10}, Hit: true})

	// Print board
	board.PrintBoard(boats, allShots)
}
