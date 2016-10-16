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
	server    *httptest.Server
	reader    io.Reader
	usersUrl  string
	imagesUrl string
)

func init() {
	server = httptest.NewServer(Handlers())

	usersUrl = fmt.Sprintf("%s/users", server.URL)
	imagesUrl = fmt.Sprintf("%s/img/api/v2.0/images", server.URL)
}

func TestCreateUser(t *testing.T) {
	userJson := `{"username": "dennis", "balance": 200}`

	reader = strings.NewReader(userJson)

	request, err := http.NewRequest("POST", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUniqueUsername(t *testing.T) {
	userJson := `{"username": "dennis", "balance": 200}`

	reader = strings.NewReader(userJson)

	request, err := http.NewRequest("POST", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 400 {
		t.Error("Bad Request expected: %d", res.StatusCode)
	}
}

func TestListUsers(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", usersUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

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
