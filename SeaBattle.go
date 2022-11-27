package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IP struct {
	ip   string
	port uint16
}

func SplitIpAndPort(str string) (string, uint16) {
	split := strings.Split(str, ":")
	ip, port := split[0], split[1]

	ui16, err := strconv.ParseUint(port, 10, 64)
	ui := uint16(ui16)

	if err != nil {
		panic(err)
	}

	return ip, ui

}

func addAlias(aliases *map[string]IP, ip string, username string) {
	realIp, port := SplitIpAndPort(ip)
	ipStruct := IP{
		ip:   realIp,
		port: port,
	}
	(*aliases)[username] = ipStruct
}

func displayAliases(aliases *map[string]IP) {
	for key, value := range *aliases {
		fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
	}
}

func displayAlias(aliases *map[string]IP, username string) {
	for key, value := range *aliases {
		if key == username {
			fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
		}
	}
}

func removeAlias(aliases *map[string]IP, username string) {
	for key, _ := range *aliases {
		if key == username {
			delete(*aliases, username)
			fmt.Println(username + " has been deleted.")
		}
	}
}

func getIpOff(username string, aliases *map[string]IP) (string, uint16) {
	for key, value := range *aliases {
		if key == username {
			return value.ip, value.port
		}
	}
	return "", 0
}

/*
func testAliases(aliases *map[string]IP) {
	addAlias(aliases, "192.168.0.1:55542", "Noam")
	i, p := getIpOff("Noam", aliases)
	fmt.Printf("%s:%d\n", i, p)
	displayAliases(aliases)
	displayAlias(aliases, "Noam")
	removeAlias(aliases, "Noam")
}
*/

func main() {
	//aliases := make(map[string]IP)
	//testAliases(&aliases)

}
