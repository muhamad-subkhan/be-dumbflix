package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func EpisodeRoutes(r *mux.Router) {
	episodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(episodeRepository)

	r.HandleFunc("/episode", middleware.Auth(h.CreateEpisode)).Methods("POST") // Add
	r.HandleFunc("/episodes", h.FindEpisode).Methods("GET") // get all
	r.HandleFunc("/episode/{id}", h.GetEpisode).Methods("GET") // get episode by id
	r.HandleFunc("/episode/{id}", middleware.Auth(h.UpdateEpisode)).Methods("PATCH") //update episode
	r.HandleFunc("/episode/{id}", middleware.Auth(h.DeleteEpisode)).Methods("DELETE") // delete
}
