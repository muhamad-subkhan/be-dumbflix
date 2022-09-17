package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/category", middleware.Auth( h.CreateCategory)).Methods("POST") // ADD
	r.HandleFunc("/categories", h.Findcategories).Methods("GET") //Get All
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET") // Get category
	r.HandleFunc("/category/{id}", middleware.Auth(h.UpdateCategory)).Methods("PATCH") // update
	r.HandleFunc("/category/{id}", middleware.Auth(h.DeleteCategory)).Methods("DELETE") // Delete

}