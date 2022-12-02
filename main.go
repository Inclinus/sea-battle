package main

import (
	"fmt"
	"internal/board"
	"internal/ip"
)

func main() {
	ip, port := ip.SplitIpAndPort("192.168.1.1:80")
	fmt.Printf("IP: %s\nPort: %d\n", ip, port)
}
