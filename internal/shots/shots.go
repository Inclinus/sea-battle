package shots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sea-battle/internal/boats"
	"sea-battle/internal/ip"
	"sea-battle/internal/utils"
	"strconv"
	"time"
)

type Shot struct {
	Position utils.Position
	Hit      bool
}

var AllShots []Shot

// Function to return an array of all shots
func GetShots() []Shot {
	return AllShots
}

// Function to check if a shot is a hit or not and return a boolean
func IsShot(boats [5]boats.Boat, position utils.Position) bool {

	// Concatenate all boats' positions
	var allBoatsPositions []utils.Position
	for _, boat := range boats {
		allBoatsPositions = append(allBoatsPositions, boat.Position...)
	}

	// Check if there is a boat at this position
	for _, boatPosition := range allBoatsPositions {
		if boatPosition.X == position.X && boatPosition.Y == position.Y {
			return true
		}
	}
	return false
}

func RequestHit(clientIP ip.IP, pos utils.Position) bool {

	port := strconv.Itoa(int(clientIP.Port))
	url := "http://" + clientIP.Ip + ":" + port + "/hit"

	jsonValue, _ := json.Marshal(pos)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := client.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	//set HTTP request header Content-Type (optional)
	//req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		//fmt.Println(err)
		fmt.Println("On dirait que votre adversaire est parti, tant pis !")
		return false
	}
	defer request.Body.Close()
	body, err := io.ReadAll(request.Body)

	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		return false
	}
	result := string(body)
	//fmt.Println(bodyString)
	if result == "true" {
		fmt.Println("Touché !")
	} else {
		fmt.Println("Raté !")
	}
	return true
}

//func MainHITTEST() {
//
//	// TODO: Select an aliases instead of IP
//	var clientIP ip.IP
//	clientIP.Ip = "127.0.0.1"
//	clientIP.Port = 4567
//
//	RequestHit(clientIP, utils.Position{X: 8, Y: 7})
//}
