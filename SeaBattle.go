package main

import (
	"fmt"
	"net/http"
)

func hitHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:

	}
}

func boatsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:

	}
}

func boardHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:

	}
}

func main() {
	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/boats", boatsHandler)
	http.HandleFunc("/hit", hitHandler)
	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		fmt.Printf("ERROR OCCURRED WHILE LAUNCHING SERVER :\n%v", err)
		return
	}
}
