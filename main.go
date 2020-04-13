package main

import (
	"fmt"
	"game21/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	gameController := controller.GameController{}
	router := mux.NewRouter()
	router.HandleFunc("/ws", gameController.GameHandler()).Methods("GET")

	fmt.Println("Server is listening...")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":1010", loggedRouter))
}