package routes

import (
	"net/http"

	"go_mefo-music-server/server/controllers"

	"github.com/gorilla/mux"
)

func SongRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/songs", controllers.AddSong).Methods(http.MethodPost)
	router.HandleFunc("/api/songs", controllers.GetAllSongs).Methods(http.MethodGet)
	return router
}
