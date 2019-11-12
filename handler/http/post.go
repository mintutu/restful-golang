package handler

import (
	"net/http"

	"github.com/mintutu/restful-golang/repository"
	"github.com/mintutu/restful-golang/repository/post"
)

func NewPostHandler() *Post {
	return &Post{
		repo: post.NewRestPostRepo(),
	}
}

type Post struct {
	repo repository.PostRepo
}

// Fetch all post data
func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	posts, _ := p.repo.Fetch(r.Context())
	RespondwithJSON(w, http.StatusOK, posts)
}
