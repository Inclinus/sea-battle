package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sea-battle/internal/boats"
	"sea-battle/internal/menu"
	"sea-battle/internal/shots"
	"sea-battle/internal/utils"
	"strconv"
)

// Function to print in navigator with Fprintln
func printLnInNav(msg string, w *http.ResponseWriter) {
	_, err := fmt.Fprintln(*w, msg)
	if err != nil {
		fmt.Println("Une erreur est survenue.")
		return
	}
}

// Function to print in navigator with Fprint
func printInNav(msg string, w *http.ResponseWriter) {
	_, err := fmt.Fprint(*w, msg)
	if err != nil {
		fmt.Println("Une erreur est survenue.")
		return
	}
}

// Handle the hit request
func hitHandler(writer http.ResponseWriter, request *http.Request) {
	// GO ROUTINE REMOVED
	switch request.Method {
	case http.MethodPost:
		var pos utils.Position
		// Decode don't work in go routine WHY ?!
		err := json.NewDecoder(request.Body).Decode(&pos)

		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println(pos.X)
		//fmt.Println(pos.Y)

		boats := boats.GenerateRandomBoats()

		result := shots.IsShot(boats, pos)
		resultConverted := strconv.FormatBool(result)

		allShots := shots.GetShots()

		allShots = append(allShots, shots.Shot{Position: pos, Hit: result})

		//fmt.Println("------------------")
		//fmt.Println(result)
		//fmt.Println("------------------")

		// Return the result of the shot
		printLnInNav(resultConverted, &writer)
	default:
		printLnInNav("Bad Request", &writer)
	}
}

// Handle boats request
func boatsHandler(writer http.ResponseWriter, request *http.Request) {
	go func() {
		switch request.Method {
		case http.MethodGet:
		default:
			printLnInNav("Bad Request", &writer)
		}
	}()
}

// Handle board request
func boardHandler(writer http.ResponseWriter, request *http.Request) {
	go func() {
		switch request.Method {
		case http.MethodGet:
		default:
			printLnInNav("Bad Request", &writer)
		}
	}()
}

func launchServer() {
	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/boats", boatsHandler)
	http.HandleFunc("/hit", hitHandler)

	err := http.ListenAndServe(":4567", nil)
	if err != nil {
		fmt.Printf("ERROR OCCURRED WHILE LAUNCHING SERVER :\n%v", err)
		return
	}
}

func main() {
	go launchServer()

	menu.DisplayMenu()
}
