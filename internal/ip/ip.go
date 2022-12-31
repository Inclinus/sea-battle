package ip

import (
	"fmt"
	"strconv"
	"strings"
)

type IP struct {
	ip   string
	port uint16
}

// SplitIpAndPort This function split an ip "192.168.0.1:8080" to a string "192.168.0.1" and a port 8080 as uint16.
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

// This function add an association between a provided IP and a provided username.
func addAlias(aliases *map[string]IP, ip string, username string) {
	realIp, port := SplitIpAndPort(ip)
	ipStruct := IP{
		ip:   realIp,
		port: port,
	}
	(*aliases)[username] = ipStruct
}

// This function displays all the associations betweens IP and usernames.
func displayAliases(aliases *map[string]IP) {
	for key, value := range *aliases {
		fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
	}
}

// This function displays the associated IP of the username provided.
func displayAlias(aliases *map[string]IP, username string) {
	for key, value := range *aliases {
		if key == username {
			fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
		}
	}
}

// This function remove the associated IP of the username provided.
func removeAlias(aliases *map[string]IP, username string) {
	for key, _ := range *aliases {
		if key == username {
			delete(*aliases, username)
			fmt.Println(username + " has been deleted.")
		}
	}
}

// This function returns the IP of a provided username, returning IP and PORT.
func getIpOf(username string, aliases *map[string]IP) (string, uint16) {
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
}
