package router

import (
	"internal/interface/handler"

	"github.com/gorilla/mux"
)

func NewRouter(userHandler *handler.UserHandler, workHandler *handler.WorkHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/works", workHandler.Creatework).Methods("POST")
	r.HandleFunc("/works", workHandler.GetworkByID).Methods("GET")
	return r
}
