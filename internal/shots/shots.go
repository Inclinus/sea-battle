package shots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sea-battle/internal/boats"
	"sea-battle/internal/utils"
	"strconv"
)

type IP struct {
	ip   string
	port uint16
}

type Shot struct {
	// Player => To do in another branch
	Position utils.Position
	Hit      bool
}

var AllShots []Shot

func GetShots() []Shot {
	return AllShots
}

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

func requestHit(clientIP IP, pos utils.Position) {

	port := strconv.Itoa(int(clientIP.port))
	url := "http://" + clientIP.ip + ":" + port + "/hit"

	jsonValue, _ := json.Marshal(pos)
	request, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	// set HTTP request header Content-Type
	//req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		return
	}
	// Log the request body
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MainHITTEST() {

	// TODO: Select an aliases instead of IP
	var clientIP IP
	clientIP.ip = "127.0.0.1"
	clientIP.port = 4567

	requestHit(clientIP, utils.Position{X: 8, Y: 7})
}
