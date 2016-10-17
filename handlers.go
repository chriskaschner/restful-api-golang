package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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
	Result_Label_2 string  `json:"result_label_2"`
	Result_Score_2 float32 `json:"result_score_2"`
	Result_Label_3 string  `json:"result_label_3"`
	Result_Score_3 float32 `json:"result_score_3"`
}

// Image info
type Image struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Results Result `json:"results"`
	Resize  bool   `json:"resize"`
	Size    Size   `json:"size"`
}

type ImageParams struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Sizes []Size
type Results []Result
type Images []Image

type User struct {
	Id           uint32 `json:"id"`
	Username     string `json:"username"`
	MoneyBalance uint32 `json:"balance"`
	Title        string `json:"title"`
}

type UserParams struct {
	Username     string `json:"username"`
	MoneyBalance uint32 `json:"balance"`
	Title        string `json:"title"`
}

var images Images

var userIdCounter uint32 = 0
var imgIdCounter int = 0

var userStore = []User{}
var ImgStore = []Image{}

// func init() {
// 	imageJson := `{Title: "Nikes", Url: "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
//
// 	CreateImageHandler(`{Title: "Nikes", Url: "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`)
// 	CreateImageHandler(Image{Title: "Altras", Url: "https://s3-us-west-2.amazonaws.com/imgdirect/altra.jpg"})
// }

func CreateImageHandler(w http.ResponseWriter, r *http.Request) {
	p := ImageParams{}

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

	img := Image{
		Id:    imgIdCounter,
		Title: p.Title,
		Url:   p.Url,
	}

	ImgStore = append(ImgStore, img)

	imgIdCounter += 1

	w.WriteHeader(http.StatusCreated)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	p := UserParams{}

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

	err = validateUniqueness(p.Title)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := User{
		Id:           userIdCounter,
		Username:     p.Username,
		MoneyBalance: p.MoneyBalance,
		Title:        p.Title,
	}

	userStore = append(userStore, u)

	userIdCounter += 1

	w.WriteHeader(http.StatusCreated)
}

func validateUniqueness(title string) error {
	for _, u := range userStore {
		if u.Title == title {
			return errors.New("Title is already used")
		}
	}

	return nil
}

func ValidateUnique(url string) error {
	for _, u := range ImgStore {
		if u.Url == url {
			return errors.New("url is already used")
		}
	}

	return nil
}
func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(userStore)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(users)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "not found!")
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["Url"]
	// todo: Search for image amongst existing pages
	// if not found, return 404
	// if found return image json
	fmt.Fprintln(w, "Get Image:", imageId)
}

func Inference(w http.ResponseWriter, r *http.Request) {
	Result_Score, Result_Label := inception.Inference()
	// inf_result, err := json.Marshal(inception.Inference())
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// w.Write(inf_result)
	// strResultScore := strconv.FormatFloat(float64(Result_Score_1), 'f', 3, 64)
	// mapD := map[string]float32{"Result_Score_1": Result_Score_1}
	// mapD := map[string]string{"Result_Score_1": strResultScore, "Result_Label_1": Result_Label_1}
	// mapB, _ := json.Marshal(mapD)
	// w.Write(mapB)
	InfRes := &Result{
		Result_Label_1: Result_Label,
		Result_Score_1: Result_Score,
	}
	InfBody, _ := json.Marshal(InfRes)
	w.Write(InfBody)

	// p := ImageParams{}
	//
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	//   fmt.Printf("Error: %s\n", err)
	//   w.WriteHeader(http.StatusInternalServerError)
	//   return
	// }
	//
	// for _, u := range ImgStore {
	//   if u.Url == url {
	//     return errors.New("url is already used")
	//   }
	// }
	// img := Image{
	//   Id:    imgIdCounter,
	//   Title: p.Title,
	//   Url:   p.Url,
	// }
	//
	// ImgStore = append(ImgStore, img)
	// //Inference info
	// type Result struct {
	// 	Result_Label_1 string  `json:"result_label_1"`
	// 	Result_Score_1 float32 `json:"result_score_1"`
	// 	Result_Label_2 string  `json:"result_label_2"`
	// 	Result_Score_2 float32 `json:"result_score_2"`
	// 	Result_Label_3 string  `json:"result_label_3"`
	// 	Result_Score_3 float32 `json:"result_score_3"`
	// }
	//
	// // Image info
	// type Image struct {
	// 	Id      int    `json:"id"`
	// 	Title   string `json:"title"`
	// 	Url     string `json:"url"`
	// 	Results Result `json:"results"`
	// 	Resize  bool   `json:"resize"`
	// 	Size    Size   `json:"size"`
	// }
}

func GetImageSize(w http.ResponseWriter, r *http.Request) {
	url := "http://i.imgur.com/Peq1U1u.jpg"
	height, width := ImageSize(url)
	// imgsize, err := json.Marshal(ImageSize(url))
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	mapD := map[string]int{"Height": height, "Width": width}
	mapB, _ := json.Marshal(mapD)
	w.Write(mapB)
	//#todo format and insert json into existing ImgStore
}

func ImagesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ImgStore); err != nil {
		panic(err)
	}
}
