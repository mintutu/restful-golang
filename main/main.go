package main

//https://github.com/s1s1ty/go-mysql-crud
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	ph "github.com/mintutu/restful-golang/handler/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	pHandler := ph.NewPostHandler()
	cHandler := ph.NewCommentHandler()
	r.Get("/comments", cHandler.Fetch)
	r.Get("/posts", pHandler.Fetch)
	r.Mount("/debug", middleware.Profiler())

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}
