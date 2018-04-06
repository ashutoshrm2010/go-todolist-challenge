package app

import (
	"net/http"

	"github.com/PeakActivity/go-todolist-challenge/app/lib"
)

// Message ...
type Message struct {
	Message string `json:"message"`
}

// HelloWorld ...
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}

	m := Message{"Hello User!"}
	res.SendOK(m)
}
