package main

import (
	"dumbflix/database"
	"dumbflix/pkg/mysql"
	"dumbflix/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Centent-type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode("Hello World")
	// })

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}