package route

import (
	"fmt"
	"net/http"
)

type Route struct {
	execute func(http.ResponseWriter, *http.Request)
}

type RouteController struct {
	routes map[string]Route
	error  func(http.ResponseWriter, *http.Request, int, string)
}

func NewRouteController() *RouteController {
	controller := new(RouteController)
	controller.routes = make(map[string]Route)
	controller.error = func(writer http.ResponseWriter, request *http.Request, status int, body string) {
		writer.WriteHeader(status)
		fmt.Fprintln(writer, "Error:", body)
	}
	return controller
}

func (controller *RouteController) ErrorHandler(handler func(http.ResponseWriter, *http.Request, int, string)) {
	controller.error = handler
}

func (controller RouteController) exists(path string) bool {
	if _, ok := controller.routes[path]; ok {
		return true
	}
	return false
}

func (controller RouteController) get(path string) Route {
	return controller.routes[path]
}

func (controller RouteController) register(path string, route Route) {
	controller.routes[path] = route
}

func (controller RouteController) Register(path string, f func(writer http.ResponseWriter, request *http.Request)) {
	controller.register(path, Route{f})
}

func (controller RouteController) Handle(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Path
	if controller.exists(url) {
		controller.get(url).execute(writer, request)
	} else {
		controller.error(writer, request, 404, "Route not found")
	}
}
