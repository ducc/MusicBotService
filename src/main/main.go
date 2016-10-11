package main

import (
    "net/http"
    "fmt"
    "encoding/json"
    "../route"
)

const API_VERSION = "1"

type abstractResponse struct {
    Error   bool        `json:"error"`
    Version string      `json:"version"`
    Content interface{} `json:"content"`
}

func main() {
    controller := route.NewRouteController()
    controller.ErrorHandler(errorHandler)
    registerRoutes(controller)
    http.HandleFunc("/", controller.Handle)
    http.ListenAndServe(":8080", nil)
    fmt.Println("Server started!")
}

func registerRoutes(controller *route.RouteController) {
    controller.Register("/", indexRoute)
}

func errorHandler(writer http.ResponseWriter, _ *http.Request, status int, body string) {
    response := abstractResponse{true, API_VERSION, body}
    writer.WriteHeader(status)
    json.NewEncoder(writer).Encode(response)
}

func indexRoute(writer http.ResponseWriter, _ *http.Request) {
    json.NewEncoder(writer).Encode(abstractResponse{false, API_VERSION, "Hello, world!"})
}