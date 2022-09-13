package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepositories := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepositories)

	r.HandleFunc("/Film", h.CreateFilm).Methods("POST") //add film
	r.HandleFunc("/Film/{id}", h.GetFilm).Methods("GET") //Get Film
	r.HandleFunc("/Film", h.FindFilm).Methods("GET") // Get All
	r.HandleFunc("/Film/{id}", h.UpdateFilm).Methods("PATCH") // Update
	r.HandleFunc("/Film/{id}", h.DeleteFilm).Methods("DELETE") //Delete
}


