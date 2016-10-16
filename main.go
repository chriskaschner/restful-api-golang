package main

import (
	"fmt"
	"net/http"

	"github.com/chriskaschner/restful-api-golang/api"
)

func main() {
	fmt.Println("Server starting")
	http.ListenAndServe(":8080", api.Handlers())
}
