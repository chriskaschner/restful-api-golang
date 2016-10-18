package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server         *httptest.Server
	reader         io.Reader
	usersUrl       string
	imagesUrl      string
	InferenceUrl   string
	IndImageUrl    string
	BadIndImageUrl string
)

func init() {
	server = httptest.NewServer(Handlers())

	usersUrl = fmt.Sprintf("%s/users", server.URL)
	imagesUrl = fmt.Sprintf("%s/img/api/v2.0/images", server.URL)
	InferenceUrl = fmt.Sprintf("%s/img/api/v2.0/inference", server.URL)
	IndImageUrl = fmt.Sprintf("%s/img/api/v2.0/images/0", server.URL)
	BadIndImageUrl = fmt.Sprintf("%s/img/api/v2.0/images/99", server.URL)

}

// func TestCreateUser(t *testing.T) {
// 	userJson := `{"username": "dennis", "balance": 200}`
//
// 	reader = strings.NewReader(userJson)
//
// 	request, err := http.NewRequest("POST", usersUrl, reader)
//
// 	res, err := http.DefaultClient.Do(request)
//
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	if res.StatusCode != 201 {
// 		t.Errorf("Success expected: %d", res.StatusCode)
// 	}
// }
func TestCreateImage(t *testing.T) {
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`

	reader = strings.NewReader(ImageJson)

	request, err := http.NewRequest("POST", imagesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUniqueImage(t *testing.T) {
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`

	reader = strings.NewReader(ImageJson)

	request, err := http.NewRequest("POST", imagesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 400 {
		t.Errorf("Bad Request expected: %d", res.StatusCode)
	}
}

// func TestUniqueUsername(t *testing.T) {
// 	userJson := `{"username": "dennis", "balance": 200}`
//
// 	reader = strings.NewReader(userJson)
//
// 	request, err := http.NewRequest("POST", usersUrl, reader)
//
// 	res, err := http.DefaultClient.Do(request)
//
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	if res.StatusCode != 400 {
// 		t.Errorf("Bad Request expected: %d", res.StatusCode)
// 	}
// }
func TestListImages(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", imagesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Get List of Images- received %d expected 200", res.StatusCode)
	}
}

func TestIndividualImageGood(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", IndImageUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Get Individual Image- received %d expected 200", res.StatusCode)
	}
}

func TestIndividualImageBad(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", BadIndImageUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 404 {
		t.Errorf("Get Individual Image- received %d expected 404", res.StatusCode)
	}
}

func TestImageInference(t *testing.T) {
	// setup image in Image Store
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
	reader = strings.NewReader(ImageJson)
	request, err := http.NewRequest("POST", imagesUrl, reader)
	res, err := http.DefaultClient.Do(request)

	reader = strings.NewReader("")
	InferenceUrl = InferenceUrl + "/0"
	request, err = http.NewRequest("GET", InferenceUrl, reader)
	res, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Bad Request expected: %d", res.StatusCode)
	}
}

func TestBadImageInference(t *testing.T) {
	// setup image in Image Store
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
	reader = strings.NewReader(ImageJson)
	request, err := http.NewRequest("POST", imagesUrl, reader)
	res, err := http.DefaultClient.Do(request)

	BlankJson := ""
	reader = strings.NewReader(BlankJson)
	InferenceUrl = InferenceUrl + "/999"
	request, err = http.NewRequest("GET", InferenceUrl, reader)
	res, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("Bad Request expected: %d", res.StatusCode)
	}
}
