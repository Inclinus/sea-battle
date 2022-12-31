package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var Aliases map[string]IP

type IP struct {
	ip   string
	port uint16
}

type User struct {
	Username string
	Ip       IP
}

// SplitIpAndPort This function split an ip "192.168.0.1:8080" to an IP struct with : as IP "192.168.0.1" as string and as PORT 8080 as uint16.
func SplitIpAndPort(str string) IP {
	split := strings.Split(str, ":")
	ip, port := split[0], split[1]

	ui16, err := strconv.ParseUint(port, 10, 64)
	ui := uint16(ui16)

	if err != nil {
		panic(err)
	}

	return IP{ip: ip, port: ui}
}

// This function add an association between a provided IP and a provided username.
func addAlias(ip string, username string) {
	(Aliases)[username] = SplitIpAndPort(ip)
}

// This function displays all the associations betweens IP and usernames.
func displayAliases() {
	for key, value := range Aliases {
		fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
	}
}

// This function displays the associated IP of the username provided.
func displayAlias(username string) {
	for key, value := range Aliases {
		if key == username {
			fmt.Printf("%s (%s:%d)\n", key, value.ip, value.port)
		}
	}
}

// This function remove the associated IP of the username provided.
func removeAlias(username string) {
	for key, _ := range Aliases {
		if key == username {
			delete(Aliases, username)
			fmt.Println(username + " has been deleted.")
		}
	}
}

// This function returns the IP of a provided username, returning IP and PORT.
func getIpOf(username string) (string, uint16) {
	for key, value := range Aliases {
		if key == username {
			return value.ip, value.port
		}
	}
	return "", 0
}

// This function allows to store every alias in a json file
func SaveAlias() {
	var userList []User
	for key, value := range Aliases {
		userList = append(userList, User{Username: key, Ip: IP{ip: value.ip, port: value.port}})
	}
	finalJson, err := json.MarshalIndent(userList, "", "")
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("alias.json", finalJson, 0644)
}

func ReceiveAlias() {
	var users []User
	file, _ := os.ReadFile("alias.json")
	_ = json.Unmarshal(file, &users)
	for indexUser := range users {
		(Aliases)[users[indexUser].Username] = users[indexUser].Ip
	}
}

func GetAlias() map[string]IP {
	return Aliases
}

func InitAliases() {
	Aliases = make(map[string]IP)
	ReceiveAlias()
}
