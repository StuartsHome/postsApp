package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
	"github.com/gorilla/mux "
	"./service"
	"./repository"
)

var (
	postRepository repository.PostRepository = repository.NewFilestoreRepository()
	postService service.PostService = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postContoller.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)
	httpRouter.SERVE(port)

}
