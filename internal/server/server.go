package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sea-battle/internal/board"
	"sea-battle/internal/boats"
	"sea-battle/internal/stats"
	"sea-battle/internal/utils"
	"strconv"
)

// Function to print in navigator with Fprintln
func printLnInNav(msg string, w *http.ResponseWriter) {
	_, err := fmt.Fprintln(*w, msg)
	if err != nil {
		fmt.Println("Une erreur est survenue. : " + err.Error())
		return
	}
}

// Function to print in navigator with Fprint
func printInNav(msg string, w *http.ResponseWriter) {
	_, err := fmt.Fprint(*w, msg)
	if err != nil {
		fmt.Println("Une erreur est survenue : " + err.Error())
		return
	}
}

// Handle board request
func pingHandler(writer http.ResponseWriter, request *http.Request) {
	// SAME PROBLEM AS HIT HANDLER
	switch request.Method {
	case http.MethodGet:
		printInNav("pong", &writer)
	default:
		printLnInNav("Bad Request", &writer)
	}
}

// Handle the hit request
func hitHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		var pos utils.Position

		err := json.NewDecoder(request.Body).Decode(&pos)

		if err != nil {
			fmt.Println(err)
			return
		}

		result := board.AddShot(pos)

		resultConverted := strconv.FormatBool(result)
		printLnInNav(resultConverted, &writer)

		// Check if all boats are destroyed
		if boats.GetAliveBoats(*board.GetBoatsBoard()) == 0 {
			// Update stats
			stats.AddGameLost()

			// If all boats are destroyed, print a message and ask to restart a game
			fmt.Print("\nTous vos bateaux ont été coulés ! Vous avez perdu !\n")
			fmt.Println("Appuyez sur Entrée pour relancer une partie...")
			fmt.Scanln()

			// Restart the game by asking the player to place their boats
			var boatsBoard [5]boats.Boat
			choice := "CHOICE"
			for choice != "O" {
				boatsBoard = boats.GenerateRandomBoats()
				board.PrintBoard(boatsBoard, false)
				fmt.Println("Voici votre board, est-ce qu'il vous satisfait ? (O/N)")
				fmt.Scanf("%s\n", &choice)
			}
			board.InitBoatsBoard(boatsBoard)
		}

	default:
		printLnInNav("Bad Request", &writer)
	}
}

// Handle boats request
func boatsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		printInNav(strconv.Itoa(int(boats.GetAliveBoats(*board.GetBoatsBoard()))), &writer)

	default:
		printInNav("Bad Request", &writer)
	}
}

// Handle board request
func boardHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:

		result := board.PrintBoard2(*board.GetBoatsBoard(), true)
		printInNav(result, &writer)

	default:
		printInNav("Bad Request", &writer)
	}
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
