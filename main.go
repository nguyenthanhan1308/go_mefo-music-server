package main

import (
	"fmt"
	"go_mefo-music-server/server/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	r := routes.SongRoutes()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is mefo-music.vercel.app api using GO")
	}).Methods(http.MethodGet)

	log.Println("Listening ...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
