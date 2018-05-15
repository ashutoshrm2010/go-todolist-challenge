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
	r.Methods("POST").Path("/item/insert").HandlerFunc(InsertItem)
	r.Methods("POST").Path("/item/edit").HandlerFunc(UpdateItem)
	r.Methods("POST").Path("/item/list").HandlerFunc(ListItem)
	r.Methods("POST").Path("/item/delete").HandlerFunc(DeleteItem)

	return r
}
