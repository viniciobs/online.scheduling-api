package main

import (
	"github.com/gorilla/mux"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/sarulabs/di"
)

func ConfigureRouter(ctn di.Container) *mux.Router {
	r := mux.NewRouter()
	request, _ := ctn.SubContainer()

	configureUserRoutes(r, request)

	return r
}

func configureUserRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get("user-handler").(*api.UsersHandler)

	r.HandleFunc("/api/users", handler.GetAll).Methods("GET")
	r.HandleFunc("/api/users/{id}", handler.GetById).Methods("GET")
	r.HandleFunc("/api/users", handler.Create).Methods("POST")
	r.HandleFunc("/api/users/{id}/activate", handler.Activate).Methods("PATCH")
	r.HandleFunc("/api/users/{id}", handler.Delete).Methods("DELETE")
}
