package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Username string
	Ip       string
	Port     uint16
}

// This function allows to store every alias in a json file
func SaveAlias(aliases *map[string]IP) {
	user := []User{}
	for key, value := range *aliases {
		user_1 := User{Username: key, Ip: value.ip, Port: value.port}
		user = append(user, user_1)

		//package this data as json data
		finalJson, err := json.MarshalIndent(user, "", "")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(finalJson))
		//fmt.Println(user)
		_ = ioutil.WriteFile("users_list.json", finalJson, 0644)
	}

}

func main() {

	aliases := make(map[string]IP)
	//addAlias(&aliases, "192.18.18.8:8080", "charbel")
	//addAlias(&aliases, "192.3.21.2:4567", "dfhfhel")
	//addAlias(&aliases, "192.3.21.2:1234", "qwerty")
	SaveAlias(&aliases)

}
