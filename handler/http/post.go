package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mintutu/restful-golang/model"
	"github.com/mintutu/restful-golang/repository"
)

func NewPostHandler(repo repository.PostRepo) *Post {
	return &Post{
		repo,
		// repo: post.NewRestPostRepo(),
		// repo: post.NewMySQLPostRepo(db.SQL),
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

// Create a new post
func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	post := model.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.Create(r.Context(), &post)
	fmt.Println(newID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}
