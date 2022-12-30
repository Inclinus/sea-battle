package main

import (
	"sea-battle/internal/menu"
	"fmt"
	"internal/ip"
	"net/http"
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
	switch request.Method {
	case http.MethodPost:
	default:
		printLnInNav("Bad Request", &writer)
	}
}

// Handle boats request
func boatsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
	default:
		printLnInNav("Bad Request", &writer)
	}
}

// Handle board request
func boardHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
	default:
		printLnInNav("Bad Request", &writer)
	}
}

func main() {
	menu.DisplayMenu()
	ip, port := ip.SplitIpAndPort("192.168.1.1:80")
	fmt.Printf("IP: %s\nPort: %d\n", ip, port)

	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/boats", boatsHandler)

	http.HandleFunc("/hit", hitHandler)

	err := http.ListenAndServe(":4567", nil)
	if err != nil {
		fmt.Printf("ERROR OCCURRED WHILE LAUNCHING SERVER :\n%v", err)
		return
	}
}
