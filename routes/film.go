package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepositories := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepositories)

	r.HandleFunc("/Film", middleware.Auth(h.CreateFilm)).Methods("POST") //add film
	r.HandleFunc("/Film/{id}", h.GetFilm).Methods("GET") //Get Film
	r.HandleFunc("/Film", h.FindFilm).Methods("GET") // Get All
	r.HandleFunc("/Film/{id}", middleware.Auth(h.UpdateFilm)).Methods("PATCH") // Update
	r.HandleFunc("/Film/{id}", middleware.Auth(h.DeleteFilm)).Methods("DELETE") //Delete
}


