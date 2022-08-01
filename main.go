package main

import (
	"fmt"
	"log"
	"net/http"

	"go_mefo-music-server/server/routes"
)

func main() {
	port := ":3333"
	r := routes.SongRoutes()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is mefo-music.vercel.app api using GO")
	}).Methods(http.MethodGet)

	log.Println("Listening ...")
	log.Fatal(http.ListenAndServe(port, r))
}
