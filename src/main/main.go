package main

import (
	"../database"
	"../route"
	"encoding/json"
	"fmt"
	"net/http"
)

const API_VERSION = "1"

var (
	conf *config
	db   *database.Database
)

type abstractResponse struct {
	Error   bool        `json:"error"`
	Version string      `json:"version"`
	Content interface{} `json:"content"`
}

func main() {
	conf = load("config.json")
	loadStatements("sql")
	var err error
	db, err = createDatabase().Open()
	if err != nil {
		fmt.Println("Error opening database connection,", err)
		return
	}
	err = createTables()
	if err != nil {
		fmt.Println("Error creating tables,", err)
		return
	}
	controller := route.NewRouteController()
	controller.ErrorHandler(errorHandler)
	controller.ApiVersion("v1")
	registerRoutes(controller)
	http.HandleFunc("/", controller.Handle)
	http.ListenAndServe(":8080", nil)
}

func errorHandler(writer http.ResponseWriter, _ *http.Request, status int, body string) {
	response := abstractResponse{true, API_VERSION, body}
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(response)
}

func registerRoutes(controller *route.RouteController) {
	controller.Register("/", "GET", indexRoute)
	controller.Register("/search/youtube", "GET", youtubeSearchRoute)
}
