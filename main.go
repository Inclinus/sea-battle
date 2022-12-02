package main

import (
	"fmt"
	"sea-battle/internal/board"
	"sea-battle/internal/boats"
	"sea-battle/internal/ip"
)

func main() {
	ip, port := ip.SplitIpAndPort("192.168.1.1:80")
	fmt.Printf("IP: %s\nPort: %d\n", ip, port)

	board.PrintEmptyBoard()

	fmt.Println(boats.GenerateRandomBoats())
}
