package main

import (
	"sea-battle/internal/ip"
	"sea-battle/internal/menu"
)

func main() {
	ip.InitAliases()
	menu.DisplayMenu()
}
