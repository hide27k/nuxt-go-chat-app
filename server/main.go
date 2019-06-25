package main

import (
	"net/http"

	"github.com/hideUW/nuxt-go-chat-app/server/infra/router"
)

func main() {
	// 静的ファイルが置いてある場所をentrypointとする。ビルドされたNuxtのファイル。
	entrypoint := "../client/nuxt-go-chat-app/dist/index.html"
	// HandlerFunc型にキャストすることで、HandlerFunc型のServeHTTPメソッドを付与する。
	router.Router.Path("/").HandlerFunc(ServeStaticFile(entrypoint))
	// 特定のPrefixを除く。
	router.Router.PathPrefix("/_nuxt/").Handler(
		http.StripPrefix("/_nuxt/", http.FileServer(http.Dir("../client/nuxt-go-chat-app/dist/_nuxt/"))))

	// 8080でサーバの立ち上げ。
	if err := http.ListenAndServe(":8080", router.Router); err != nil {
		panic(err.Error())
	}
}

// ServeStaticFile delivers static files.
// インターフェイスを用いている。
func ServeStaticFile(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}
