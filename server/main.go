package main

import (
	"net/http"

	"github.com/hideUW/nuxt-go-chat-app/server/infra/router"
)

func main() {
	// For static file
	entrypoint := "../client/nuxt-go-chat-app/dist/index.html"
	router.Router.Path("/").HandlerFunc(ServeStaticFile(entrypoint))
	router.Router.PathPrefix("/_nuxt/").Handler(
		http.StripPrefix("/_nuxt/", http.FileServer(http.Dir("../client/nuxt-go-chat-app/dist/_nuxt/"))))

	if err := http.ListenAndServe(":8080", router.Router); err != nil {
		panic(err.Error())
	}
}

// ServeStaticFile delivers static files
func ServeStaticFile(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}
