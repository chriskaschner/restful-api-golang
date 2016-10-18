package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	inception "github.com/chriskaschner/Inception-Retraining-Golang"
	"github.com/gorilla/mux"
)

type Size struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

//Inference info
type Result struct {
	Result_Label_1 string  `json:"result_label_1"`
	Result_Score_1 float32 `json:"result_score_1"`
}

// Image info
type Image struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Uri     string `json:"uri"`
	Results Result `json:"results"`
	Resize  bool   `json:"resize"`
	Size    Size   `json:"size"`
}

type Sizes []Size
type Results []Result
type Images []Image

var images Images

var imgIdCounter int = 0

var ImgStore = []Image{}

// func init() {
//
// // CreateImageHandler(`{Title: "Nikes", Url: "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`)
// // CreateImageHandler(`{Title: "Altras", Url: "https://s3-us-west-2.amazonaws.com/imgdirect/altra.jpg"}`)
// }
func CreateImageHandler(w http.ResponseWriter, r *http.Request) {
	p := Image{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = ValidateUnique(p.Url)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		preexisting, err := json.Marshal(`{error: "image already exists in DB"}`)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(preexisting)
		return
	}

	// #todo assemble Uri that includes full & raw path
	img := Image{
		Id:    imgIdCounter,
		Title: p.Title,
		Url:   p.Url,
		Uri:   r.URL.String() + "/" + strconv.Itoa(imgIdCounter),
	}

	ImgStore = append(ImgStore, img)

	imgIdCounter += 1

	w.WriteHeader(http.StatusCreated)
}

func ValidateUnique(url string) error {
	for _, u := range ImgStore {
		if u.Url == url {
			return errors.New("url is already used")
		}
	}

	return nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ImagesIndex(w http.ResponseWriter, r *http.Request) {
	ImgStoreBody, _ := json.Marshal(ImgStore)
	if len(ImgStore) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(ImgStoreBody)
	if err := json.NewEncoder(w).Encode(ImgStore); err != nil {
		panic(err)
	}
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	// retrieve image id from URL of request
	vars := mux.Vars(r)
	imageId := vars["id"]

	// Search for image amongst existing records in store
	for _, u := range ImgStore {
		IdString := strconv.Itoa(u.Id)
		// if found, return JSON for that image
		if IdString == imageId {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			ImageBody, _ := json.Marshal(u)
			w.Write(ImageBody)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// #todo, currently overwrites ALL fields
func UpdateImage(w http.ResponseWriter, r *http.Request) {
	// Read image id from url
	vars := mux.Vars(r)
	imageId := vars["id"]

	// read JSON in body
	p := Image{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Search for image amongst existing records in store
	for _, u := range ImgStore {
		IdString := strconv.Itoa(u.Id)
		// if found, update JSON for that image
		if IdString == imageId {
			u = Image{
				Title: p.Title,
				Url:   p.Url,
				// Uri:    p.Uri,
				Results: p.Results,
				Resize:  p.Resize,
				Size:    p.Size,
			}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			ImgStore[u.Id] = u
			ImageBody, _ := json.Marshal(u)
			w.Write(ImageBody)
			return
		}
		// if not found, return 404
		http.Error(w, "image not found", http.StatusNotFound)
	}
}

func RunInference(w http.ResponseWriter, r *http.Request) {
	// Read image id from url
	vars := mux.Vars(r)
	imageId := vars["id"]

	// Search for image amongst existing records in store
	for _, u := range ImgStore {
		IdString := strconv.Itoa(u.Id)
		// if found, add inference JSON to that image
		if IdString == imageId {
			Result_Score, Result_Label := inception.Inference(u.Url)

			u.Results = Result{
				Result_Label_1: Result_Label,
				Result_Score_1: Result_Score,
			}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			ImgStore[u.Id] = u
			ImageBody, _ := json.Marshal(u)
			w.Write(ImageBody)
			return
		}
		// if not found, return 404
		http.Error(w, "image not found", http.StatusNotFound)
	}
}

func GetImageSize(w http.ResponseWriter, r *http.Request) {
	// Read image id from url
	vars := mux.Vars(r)
	imageId := vars["id"]

	// Search for image amongst existing records in store
	for _, u := range ImgStore {
		IdString := strconv.Itoa(u.Id)
		// if found, add size JSON to that image
		if IdString == imageId {
			url := "http://i.imgur.com/Peq1U1u.jpg"
			height, width := ImageSize(url)

			u.Size = Size{
				Height: height,
				Width:  width,
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)

			ImgStore[u.Id] = u
			ResizeBody, _ := json.Marshal(u)
			w.Write(ResizeBody)
		}
		// if not found, return 404
		http.Error(w, "image not found", http.StatusNotFound)
	}
}
