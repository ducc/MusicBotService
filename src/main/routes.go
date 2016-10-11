package main

import (
	"encoding/json"
	"net/http"
)

func indexRoute(writer http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(writer).Encode(abstractResponse{false, API_VERSION, "Hello, world!"})
}
