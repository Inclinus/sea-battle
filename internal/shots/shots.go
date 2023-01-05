package shots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sea-battle/internal/board"
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

func GetAllShots() *[]Shot {
	return &AllShots
}

var AllShots []Shot

func AddShot(position utils.Position) bool {
	isShot := checkShot(position)

	actualShot := Shot{Position: position, Hit: isShot}

	AllShots = append(AllShots, actualShot)

	if isShot {
		checkDestroyed(getBoatAt(position))
	}

	return actualShot.Hit
}

func getBoatAt(position utils.Position) *boats.Boat {
	for _, boat := range board.GetBoatsBoard() {
		for _, pos := range boat.Position {
			if pos.X == position.X && pos.Y == position.Y {
				return &boat
			}
		}
	}
	panic("POSITION DOES NOT CORRESPOND TO A BOAT")
}

func checkDestroyed(boat *boats.Boat) {
	count := boat.Size
	for _, pos := range boat.Position {
		for _, shot := range AllShots {
			if pos.X == shot.Position.X && pos.Y == shot.Position.Y {
				count--
			}
		}
	}
	if count <= 0 {
		boat.Destroyed = true
	}
}

// Function to check if a shot is a hit or not and return a boolean
func checkShot(position utils.Position) bool {

	// Concatenate all boats' positions
	var allBoatsPositions []utils.Position
	for _, boat := range board.GetBoatsBoard() {
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
	if result == "true\n" {
		fmt.Println("Touché !")
	} else {
		fmt.Println("Raté !")
	}
	return true
}
