package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Websocket Endpoint")
}
