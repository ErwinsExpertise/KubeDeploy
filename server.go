package main

import (
	"log"
	"net/http"

	"github.com/ErwinsExpertise/KubeDeploy/handlers"
	"github.com/gorilla/mux"
)

func startServer() {
	//Create new mux router
	rout := mux.NewRouter()
	port := ":9000"

	//Endpoints for dashboard
	rout.HandleFunc("/", handlers.HomeHandler)
	rout.HandleFunc("/add", handlers.NewSiteHandler)
	rout.HandleFunc("/list", handlers.ListSiteHandler)

	rout.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("html/css/"))))

	log.Println("Listening on " + port)
	//Start listening
	log.Fatal(http.ListenAndServe(port, rout))
}
