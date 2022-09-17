package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
  userRepository := repositories.RepositoryUser(mysql.DB)
  h := handlers.HandlerUser(userRepository)

  // r.HandleFunc("/user", h.CreateUser).Methods("POST") // Create
  r.HandleFunc("/users", h.FindUsers).Methods("GET") // Get all
  r.HandleFunc("/user/{id}", h.GetUser).Methods("GET") // Get User
  r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH") // Get Update 
  r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE") // Delete
}
