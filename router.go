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
	configureModalityRoutes(r, request)
	configureUserModalityRoutes(r, request)

	return r
}

func configureUserRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get("user-handler").(*api.UsersHandler)

	r.HandleFunc("/api/users", handler.Get).Methods("GET")
	r.HandleFunc("/api/users/{id}", handler.GetById).Methods("GET")
	r.HandleFunc("/api/users", handler.Create).Methods("POST")
	r.HandleFunc("/api/users/{id}/activate", handler.Activate).Methods("PATCH")
	r.HandleFunc("/api/users/{id}/edit", handler.Edit).Methods("PATCH")
	r.HandleFunc("/api/users/{id}", handler.Delete).Methods("DELETE")
}

func configureModalityRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get("modality-handler").(*api.ModalityHandler)

	r.HandleFunc("/api/modalities", handler.Get).Methods("GET")
	r.HandleFunc("/api/modalities/{id}", handler.GetById).Methods("GET")
	r.HandleFunc("/api/modalities", handler.Create).Methods("POST")
	r.HandleFunc("/api/modalities/{id}/edit", handler.Edit).Methods("PATCH")
	r.HandleFunc("/api/modalities/{id}", handler.Delete).Methods("DELETE")
}

func configureUserModalityRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get("user-modalities-handler").(*api.UserModalitiesHandler)

	r.HandleFunc("/api/users/{id}/modalities", handler.Edit).Methods("PUT")
}
