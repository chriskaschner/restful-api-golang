package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Not Found",
		"GET",
		"/{[a-z0-9]+}",
		NotFound,
	},
	Route{
		"ImagesIndex",
		"GET",
		"/api/v2.0/imgs",
		ImagesIndex,
	},
	Route{
		"GetImage",
		"GET",
		"/api/v2.0/imgs/{ImgId}",
		RepoGetImage,
	},
	Route{
		"CreateImage",
		"POST",
		"/api/v2.0/imgs",
		CreateImage,
	},
}
