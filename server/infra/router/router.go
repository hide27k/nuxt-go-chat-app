package router

import (
	"github.com/gorilla/mux"
)

// Router is the gorilla router for API.
var Router *mux.Router

func init() {
	// Instantiation for gorilla Router.
	// NewRouter関数で*mux.Router型の変数を作る。
	r := mux.NewRouter()

	Router = r
}
