package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	TransactionRepository := repositories.RepositoryforTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST") //Create
	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET") //Get All
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET") //Get User
	r.HandleFunc("/transaction/{id}", h.UpdateTransaction).Methods("PATCH") // Updates 
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE") // Delete

}
