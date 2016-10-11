package route

import (
	"fmt"
	"net/http"
)

type Route struct {
    method string
	execute func(http.ResponseWriter, *http.Request)
}

type RouteController struct {
	routes map[string]Route
	error  func(http.ResponseWriter, *http.Request, int, string)
    apiVersion string
}

func NewRouteController() *RouteController {
	controller := new(RouteController)
	controller.routes = make(map[string]Route)
	controller.error = func(writer http.ResponseWriter, request *http.Request, status int, body string) {
		writer.WriteHeader(status)
		fmt.Fprintln(writer, "Error:", body)
	}
    controller.apiVersion = "DEFAULT_VERSION"
	return controller
}

func (controller *RouteController) ErrorHandler(handler func(http.ResponseWriter, *http.Request, int, string)) {
	controller.error = handler
}

func (controller *RouteController) ApiVersion(apiVersion string) {
    controller.apiVersion = apiVersion
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
	controller.routes["/" + controller.apiVersion + path] = route
}

func (controller RouteController) Register(path string, method string, f func(writer http.ResponseWriter, request *http.Request)) {
	controller.register(path, Route{method, f})
}

func (controller RouteController) Handle(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Path
    fmt.Println(request.Method + " " + url)
	if controller.exists(url) {
        r := controller.get(url)
        if r.method != request.Method {
            return
        }
		r.execute(writer, request)
	} else {
		controller.error(writer, request, 404, "Route not found")
	}
}
