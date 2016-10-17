package main

import (
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")

	// subrouter to add prefix for all other handlers
	s := r.PathPrefix("/img/api/v2.0").Subrouter()

	// images index
	s.HandleFunc("/images", ImagesIndex).Methods("GET")

	// new images
	s.HandleFunc("/images", CreateImageHandler).Methods("POST")

	// runs inference on an image using Inception model
	s.HandleFunc("/inference", Inference).Methods("GET")

	r.HandleFunc("/users", createUserHandler).Methods("POST")

	r.HandleFunc("/users", listUsersHandler).Methods("GET")

	return r
}
