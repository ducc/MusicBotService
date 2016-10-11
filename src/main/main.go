package main

import (
	"../route"
	"encoding/json"
	"fmt"
	"net/http"
)

const API_VERSION = "1"

var conf *config

type abstractResponse struct {
	Error   bool        `json:"error"`
	Version string      `json:"version"`
	Content interface{} `json:"content"`
}

func main() {
	conf = load("config.json")
	controller := route.NewRouteController()
	controller.ErrorHandler(errorHandler)
	registerRoutes(controller)
	http.HandleFunc("/", controller.Handle)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started!")
}

func errorHandler(writer http.ResponseWriter, _ *http.Request, status int, body string) {
	response := abstractResponse{true, API_VERSION, body}
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(response)
}

func registerRoutes(controller *route.RouteController) {
	controller.Register("/", indexRoute)
	controller.Register("/search/youtube", youtubeSearchRoute)
}
