package main

import (
	"encoding/json"
	"net/http"

	"./entity"
	"./repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" : "Error getting the posts"}`))
	}
	// result, err := json.Marshal(posts)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{"error" : "Error marshaling data"}`))
	// }
	response.WriteHeader(http.StatusOK)
	//response.Write(result)
	json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" : "Error marshaling data"}`))
	}
}
