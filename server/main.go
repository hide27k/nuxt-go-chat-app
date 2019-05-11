package main

import (
	"net/http"

	"github.com/hideUW/nuxt_go_template/server/infra/router"
)

func main() {
	// For static file
	entrypoint := "../client/nuxt_go_template/dist/index.html"
	router.Router.Path("/").HandlerFunc(ServeStaticFile(entrypoint))
	router.Router.PathPrefix("/_nuxt/").Handler(
		http.StripPrefix("/_nuxt/", http.FileServer(http.Dir("../client/nuxt_go_template/dist/_nuxt/"))))

	if err := http.ListenAndServe(":8080", router.Router); err != nil {
		panic(err.Error())
	}
}

// ServeStaticFile is deliver static files
func ServeStaticFile(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}
