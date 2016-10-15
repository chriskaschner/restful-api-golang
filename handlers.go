package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "not found!")
}

func ImagesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(images); err != nil {
		panic(err)
	}
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["ImgId"]
	// todo: Search for image amongst existing pages
	// if not found, return 404
	// if found return image json
	fmt.Fprintln(w, "Get Image:", imageId)
}

func RepoGetImage(w http.ResponseWriter, r *http.Request) {
	// 	var image Image
	//
	vars := mux.Vars(r)
	imageId := vars["ImgId"]
	// Convert string to int
	imageIDint, err := strconv.Atoi(imageId)
	// 	// todo: Search for image amongst existing pages
	// 	// if not found, return 404
	// 	// if found return image json
	t := RepoFindImage(imageIDint, err)
	// 	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	// w.WriteHeader(http.StatusCreated)
	// 	if err := json.NewEncoder(w).Encode(t); err != nil {
	// 		panic(err)
	// }
	return t
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	var image Image
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &image); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateImage(image)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
