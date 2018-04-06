package app

import (
	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {

	//Create main router
	r := mux.NewRouter().StrictSlash(true)

	/**
	 * Routes
	 */
	r.Methods("GET").Path("/").HandlerFunc(HelloWorld)
	return r
}
