package handler

import (
	"net/http"

	"github.com/mintutu/restful-golang/repository"
	"github.com/mintutu/restful-golang/repository/comment"
)

func NewCommentHandler() *Comment {
	return &Comment{
		repo: comment.NewRestCommentRepo(),
	}
}

type Comment struct {
	repo repository.CommentRepo
}

// Fetch all post data
func (p *Comment) Fetch(w http.ResponseWriter, r *http.Request) {
	posts, _ := p.repo.Fetch(r.Context())
	RespondwithJSON(w, http.StatusOK, posts)
}
