package router

import (
	"github.com/gorilla/mux"
	"github.com/online.scheduling-api/constants"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/middlewares"
	"github.com/online.scheduling-api/src/models"
	"github.com/sarulabs/di"
)

var onlyAdmin = []models.Role{models.Admin}
var onlyAdminAndWorker = []models.Role{models.Admin, models.Worker}

func ConfigureRouter(ctn di.Container) *mux.Router {
	r := mux.NewRouter()
	request, _ := ctn.SubContainer()

	configureAuthRoutes(r, request)
	configureUserRoutes(r, request)
	configureModalityRoutes(r, request)
	configureScheduleRoutes(r, request)

	return r
}

func configureAuthRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get(constants.AUTH_HANDLER).(*api.AuthHandler)

	// No auth required
	r.HandleFunc("/api/sign-in", handler.SignIn).Methods("POST")
}

func configureUserRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get(constants.USER_HANDLER).(*api.UsersHandler)

	// No auth required
	r.HandleFunc("/api/users", handler.Create).Methods("POST")

	// Auth required
	r.Handle("/api/users", middlewares.EnsureAuth(handler.Get)).Methods("GET")
	r.Handle("/api/users/{id}", middlewares.EnsureAuth(handler.GetById)).Methods("GET")
	r.Handle("/api/users/{id}/edit", middlewares.EnsureAuth(handler.Edit)).Methods("PATCH")
	r.Handle("/api/users/{id}", middlewares.EnsureAuth(handler.Delete)).Methods("DELETE")

	// Admin required
	r.Handle("/api/users/{id}/activate", middlewares.EnsureRole(handler.Activate, onlyAdmin)).Methods("PATCH")

	// Only Admin or Worker
	r.Handle("/api/users/{id}/modalities", middlewares.EnsureRole(handler.EditModalities, onlyAdminAndWorker)).Methods("PUT")
}

func configureModalityRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get(constants.MODALITY_HANDLER).(*api.ModalityHandler)

	// Auth required
	r.Handle("/api/modalities", middlewares.EnsureAuth(handler.Get)).Methods("GET")
	r.Handle("/api/modalities/{id}", middlewares.EnsureAuth(handler.GetById)).Methods("GET")

	// Admin required
	r.Handle("/api/modalities", middlewares.EnsureRole(handler.Create, onlyAdmin)).Methods("POST")
	r.Handle("/api/modalities/{id}/edit", middlewares.EnsureRole(handler.Edit, onlyAdmin)).Methods("PATCH")
	r.Handle("/api/modalities/{id}", middlewares.EnsureRole(handler.Delete, onlyAdmin)).Methods("DELETE")
}

func configureScheduleRoutes(r *mux.Router, ctn di.Container) {
	handler := ctn.Get(constants.SCHEDULE_HANDLER).(*api.SchedulesHandler)

	// Auth required
	r.Handle("/api/schedules", middlewares.EnsureAuth(handler.Get)).Methods("GET")

	// Only Admin or Worker
	r.Handle("/api/schedules", middlewares.EnsureRole(handler.Create, onlyAdminAndWorker)).Methods("POST")
	r.Handle("/api/schedules", middlewares.EnsureRole(handler.Edit, onlyAdminAndWorker)).Methods("PUT")
	r.Handle("/api/schedules", middlewares.EnsureRole(handler.Delete, onlyAdminAndWorker)).Methods("DELETE")
}
