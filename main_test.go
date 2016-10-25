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
	ImagesUrl      string
	InferenceUrl   string
	IndImageUrl    string
	BadIndImageUrl string
	ResizeUrl      string
)

func init() {
	server = httptest.NewServer(Handlers())

	ImagesUrl = fmt.Sprintf("%s/img/api/v2.0/images", server.URL)
	InferenceUrl = fmt.Sprintf("%s/img/api/v2.0/inference", server.URL)
	ResizeUrl = fmt.Sprintf("%s/img/api/v2.0/resize", server.URL)

}

func TestIndex(t *testing.T) {
	request, err := http.NewRequest("GET", server.URL, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestCreateImage(t *testing.T) {
	ImageJson := `{"Title": "Altras", "Url": "https://s3-us-west-2.amazonaws.com/imgdirect/altra.jpg"}`

	reader = strings.NewReader(ImageJson)

	request, err := http.NewRequest("POST", ImagesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestListImages(t *testing.T) {
	request, err := http.NewRequest("GET", ImagesUrl, nil)
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Get List of Images- expected 200 got: %d", res.StatusCode)
	}

}

func TestCreateImageBadJson(t *testing.T) {
	BadJson := `{"abc":"abc",}`

	reader = strings.NewReader(BadJson)

	request, err := http.NewRequest("POST", ImagesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 500 {
		t.Errorf("Bad JSON- expected 500, got: %d", res.StatusCode)
	}
}

func TestUniqueImage(t *testing.T) {
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`

	reader = strings.NewReader(ImageJson)

	request, err := http.NewRequest("POST", ImagesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 400 {
		t.Errorf("Duplicate Image received- %d, expected 400", res.StatusCode)
	}
}

func TestIndividualImageGood(t *testing.T) {
	// setup image in Image Store
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
	reader = strings.NewReader(ImageJson)
	request, err := http.NewRequest("POST", ImagesUrl, reader)
	res, err := http.DefaultClient.Do(request)

	IndImageUrl := ImagesUrl + "/1"
	request, err = http.NewRequest("GET", IndImageUrl, nil)

	res, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Get Good Individual Image- received %d expected 200", res.StatusCode)
	}
}

func TestIndividualImageBad(t *testing.T) {
	BadIndImageUrl := ImagesUrl + "/999"
	request, err := http.NewRequest("GET", BadIndImageUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 404 {
		t.Errorf("Get Bad Individual Image- received %d expected 404", res.StatusCode)
	}
}

func TestImageInference(t *testing.T) {
	// setup image in Image Store
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
	reader = strings.NewReader(ImageJson)
	request, err := http.NewRequest("POST", ImagesUrl, reader)
	res, err := http.DefaultClient.Do(request)

	InferenceUrl = InferenceUrl + "/1"
	request, err = http.NewRequest("GET", InferenceUrl, nil)
	res, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Good Inference Request- Expected 200, got: %d", res.StatusCode)
	}
}

func TestBadImageInference(t *testing.T) {
	// setup image in Image Store
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
	reader = strings.NewReader(ImageJson)
	request, err := http.NewRequest("POST", ImagesUrl, reader)
	res, err := http.DefaultClient.Do(request)

	InferenceUrl = InferenceUrl + "/999"
	request, err = http.NewRequest("GET", InferenceUrl, nil)
	res, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("Bad Inference Request- Expected 404, got: %d", res.StatusCode)
	}
}

func TestImageSize(t *testing.T) {
	// setup image in Image Store
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
	reader = strings.NewReader(ImageJson)
	request, err := http.NewRequest("POST", ImagesUrl, reader)
	res, err := http.DefaultClient.Do(request)

	GoodResizeUrl := ResizeUrl + "/1"
	request, err = http.NewRequest("GET", GoodResizeUrl, nil)
	res, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Good Resize Request- expected 200 got: %d", res.StatusCode)
	}
}

func TestBadImageSize(t *testing.T) {
	// setup image in Image Store
	ImageJson := `{"Title": "Nikes", "Url": "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"}`
	reader = strings.NewReader(ImageJson)
	request, err := http.NewRequest("POST", ImagesUrl, reader)
	res, err := http.DefaultClient.Do(request)

	BadResizeUrl := ResizeUrl + "/9999"
	request, err = http.NewRequest("GET", BadResizeUrl, nil)
	res, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 404 {
		t.Errorf("Bad Resize Request- Expected 404, got: %d", res.StatusCode)
	}
}
