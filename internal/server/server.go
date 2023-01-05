package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sea-battle/internal/board"
	"sea-battle/internal/boats"
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

// Handle board request
func pingHandler(writer http.ResponseWriter, request *http.Request) {
	//go func() {
	// GO ROUTINE REMOVED
	// SAME PROBLEM AS HIT HANDLER
	switch request.Method {
	case http.MethodGet:
		printInNav("pong", &writer)
	default:
		printLnInNav("Bad Request", &writer)
	}
	//}()
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

		result := shots.IsShot(board.GetBoatsBoard(), pos)
		resultConverted := strconv.FormatBool(result)

		shots.AddShot(shots.Shot{Position: pos, Hit: result})

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
			aliveBoats := boats.GetAliveBoats(board.GetBoatsBoard())

			fmt.Println("Il reste", aliveBoats, "bateaux en vie.")
			// TODO : Return the number of alive boats in writer

		default:
			printLnInNav("Bad Request", &writer)
		}
	}()
}

// Handle board request
func boardHandler(writer http.ResponseWriter, request *http.Request) {
	//go func() {
	switch request.Method {
	case http.MethodGet:

		result := board.PrintBoard2(board.GetBoatsBoard(), true)
		printInNav(result, &writer)

	default:
		printLnInNav("Bad Request", &writer)
	}
	//}()
}

func LaunchServer() {
	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/boats", boatsHandler)
	http.HandleFunc("/hit", hitHandler)
	http.HandleFunc("/ping", pingHandler)

	err := http.ListenAndServe(":4567", nil)
	if err != nil {
		fmt.Printf("ERROR OCCURRED WHILE LAUNCHING SERVER :\n%v", err)
		return
	}
}
