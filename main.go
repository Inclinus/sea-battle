package main

import (
	"fmt"
	"net/http"
	"sea-battle/internal/menu"
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
	go func() {
		switch request.Method {
		case http.MethodPost:
		default:
			printLnInNav("Bad Request", &writer)
		}
	}()
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
