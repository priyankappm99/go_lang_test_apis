package main

import (
	"employee/models"
	"encoding/json"
	"log"
	"net/http"
)

var Articles []models.Article

func main() {
	Articles = []models.Article{
		{Title: "Helloooooooo", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}
