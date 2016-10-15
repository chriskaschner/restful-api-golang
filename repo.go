package main

import "fmt"

var currentId int

var images Images

// Give us some seed data
func init() {
	RepoCreateImage(Image{Title: "Nikes", Url: "http://imgdirect.s3-website-us-west-2.amazonaws.com/nike.jpg"})
	RepoCreateImage(Image{Title: "Altras", Url: "https://s3-us-west-2.amazonaws.com/imgdirect/altra.jpg"})
}

func RepoFindImage(id int, err error) Image {
	for _, t := range images {
		if t.Id == id {
			return t
		}
	}
	// return empty Image if not found
	return Image{}
	// fmt.Errorf("Could not find Image with id of %d", id)
}

// Original function
// func RepoFindImage(id int) Image {
// 	for _, t := range images {
// 		if t.Id == id {
// 			return t
// 		}
// 	}
// 	// return empty Image if not found
// 	return Image{}
// }

func RepoCreateImage(t Image) Image {
	currentId += 1
	t.Id = currentId
	images = append(images, t)
	return t
}

func RepoDestroyImage(id int) error {
	for i, t := range images {
		if t.Id == id {
			images = append(images[:i], images[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Image with id of %d to delete", id)
}
