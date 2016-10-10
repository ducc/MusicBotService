package main

import (
    "net/http"
    "fmt"
    "encoding/json"
)

func main() {
    controller := NewRouteController()
    controller.ErrorHandler(func(writer http.ResponseWriter, request *http.Request, status int, body string) {
        response := struct {
            Error   bool    `json:"error"`
            Status  int     `json:"status"`
            Message string  `json:"message"`
        }{true, status, body}
        json.NewEncoder(writer).Encode(response)
    })
    controller.Register("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintln(writer, "Hello, world!")
    })
    http.HandleFunc("/", controller.Handle)
    http.ListenAndServe(":8080", nil)
}
