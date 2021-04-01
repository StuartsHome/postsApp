package controller

import (
	"encoding/json"
	"net/http"
	"postsApp/service"

	"./entity"
	"./service"
)

var (
	postService service.PostService = service.NewPostService()
)

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" : "Error getting the posts"}`))
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
	// result, err := json.Marshal(posts)
	// if err != nil {
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{"error" : "Error marshaling data"}`))
	// }
	//response.Write(result)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" : "Error marshaling data"}`))
		return
	}
	err := postService.Validate(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" : "Error marshaling data"}`))
		return
	}

	postService.Create(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
