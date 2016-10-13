package main

type Image struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Results string `json:"results"`
	Resize  bool   `json:"resize"`
	Size    string `json:"size"`
}

type Images []Image
