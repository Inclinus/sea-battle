package ip

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IP struct {
	Ip   string
	Port uint16
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
func AddAlias(aliases *map[string]IP, ip string, username string) {
	realIp, port := SplitIpAndPort(ip)
	ipStruct := IP{
		Ip:   realIp,
		Port: port,
	}
	(*aliases)[username] = ipStruct
}

func isConnected(clientIP IP) bool {
	port := strconv.Itoa(int(clientIP.Port))
	url := "http://" + clientIP.Ip + ":" + port + "/ping"

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Une erreur est survenue.")
		return false
	}
	result := string(body)
	if result == "pong" {
		return true
	}
	return false
}

// This function displays all the associations betweens IP and usernames.
func DisplayAliases(aliases *map[string]IP) {
	fmt.Println("------------------------------")
	fmt.Println("Liste des aliases :")
	for key, value := range *aliases {
		var clientIP IP
		clientIP.Ip = value.Ip
		clientIP.Port = value.Port
		if isConnected(clientIP) {
			fmt.Printf("%s (%s:%d) | ✔️ Connecté \n", key, value.Ip, value.Port)
		} else {
			fmt.Printf("%s (%s:%d) | ❌ Hors-Ligne \n", key, value.Ip, value.Port)
		}
		//fmt.Printf("%s (%s:%d)\n", key, value.Ip, value.Port)
	}
	fmt.Println("------------------------------")
}

// This function displays the associated IP of the username provided.
func DisplayAlias(aliases *map[string]IP, username string) {
	for key, value := range *aliases {
		if key == username {
			fmt.Printf("%s (%s:%d)\n", key, value.Ip, value.Port)
		}
	}
}

func AliasIsExist(aliases *map[string]IP, username string) bool {
	for key := range *aliases {
		if key == username {
			return true
		}
	}
	return false
}

// This function remove the associated IP of the username provided.
func RemoveAlias(aliases *map[string]IP, username string) {
	for key, _ := range *aliases {
		if key == username {
			delete(*aliases, username)
			fmt.Println(username + " has been deleted.")
		}
	}
}

// This function returns the IP of a provided username, returning IP and PORT.
func GetIpOf(username string, aliases *map[string]IP) (string, uint16) {
	for key, value := range *aliases {
		if key == username {
			return value.Ip, value.Port
		}
	}
	return "", 0
}

func GetIpOf2(username string, aliases *map[string]IP) IP {
	for key, value := range *aliases {
		if key == username {
			var clientIP IP
			clientIP.Ip = value.Ip
			clientIP.Port = value.Port
			return clientIP
		}
	}
	return IP{}
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
