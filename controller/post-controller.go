package controller

import (
	"encoding/json"
	"net/http"
	"postsApp/service"

	"../errors"

	"./entity"
	"./service"
)

var (
	postService service.PostService = service.NewPostService()
)

type postController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
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

func AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshaling data"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
