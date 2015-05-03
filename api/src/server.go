package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shorrockin/clikr/log"
)

// defines a structure passed into the map method that tells the router
// where to send things
type Route struct {
	method  string
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

// array containing all the routing information for our requests
var routes []Route

// public method called by other classes to define different routing instructions
func Map(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	routes = append(routes, Route{method, path, handler})
}

// callback executed after ListenAndServer to set everything up
func initRouter() http.Handler {
	router := mux.NewRouter()

	for _, value := range routes {
		log.Debug("bunding %v %v", value.method, value.path)
		router.HandleFunc(value.path, value.handler).Methods(value.method)
	}

	return router
}

// main entry point for the server
func main() {
	log.Debug("clikr server booting up, binding to port 3000")
	http.ListenAndServe(":3000", initRouter())
}
