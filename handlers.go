package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ImagesIndex(w http.ResponseWriter, r *http.Request) {
	images := Images{
		Image{Title: "Nikes", Url: "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"},
		Image{Title: "Altras", Url: "https://s3-us-west-2.amazonaws.com/imgdirect/altra.jpg"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(images); err != nil {
		panic(err)
	}
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["ImgId"]
	fmt.Fprintln(w, "Get Image:", todoId)
}
