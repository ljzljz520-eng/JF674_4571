package main

import (
	"galleryline/api"
	"galleryline/service"
	"galleryline/signaling"
	"galleryline/storage"
	"log"
	"net/http"
)

func main() {
	s, e := storage.Open("galleryline.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	d := service.NewDirectory(s)
	r := signaling.NewRouter()
	c := service.NewCallManager(s, d, r)
	log.Fatal(http.ListenAndServe(":8080", api.NewServer(d, c).Handler()))
}
