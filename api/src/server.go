package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shorrockin/clikr/log"
)

func main() {
	log.Debug("clikr server booting up, binding to port 3000")
	router := mux.NewRouter()
	http.ListenAndServe(":3000", router)
}
