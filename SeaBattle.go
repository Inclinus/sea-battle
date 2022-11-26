package main

import (
	"strconv"
	"strings"
)

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
func main() {

}
