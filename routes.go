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
		"ImagesIndex",
		"GET",
		"/images",
		ImagesIndex,
	},
	Route{
		"GetImage",
		"GET",
		"/images/{ImgId}",
		GetImage,
	},
	Route{
		"CreateImage",
		"POST",
		"/images",
		CreateImage,
	},
}
