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

	// specific image info
	s.HandleFunc("/images/{id:[0-9]+}", GetImage).Methods("GET")

	// update an image
	s.HandleFunc("/images/{id:[0-9]+}", UpdateImage).Methods("PUT")

	// create new image
	s.HandleFunc("/images", CreateImageHandler).Methods("POST")

	// runs inference on an image using Inception model
	s.HandleFunc("/inference/{id:[0-9]+}", RunInference).Methods("GET")

	// gets image size
	s.HandleFunc("/resize/{id:[0-9]+}", GetImageSize).Methods("GET")

	return r
}
