package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/category", h.CreateCategory).Methods("POST") // ADD
	r.HandleFunc("/categories", h.Findcategories).Methods("GET") //Get All
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET") // Get category
	r.HandleFunc("/category/{id}", h.UpdateCategory).Methods("PATCH") // update
	r.HandleFunc("/category/{id}", h.DeleteCategory).Methods("DELETE") // Delete

}