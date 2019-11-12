package repository

import (
	"context"

	"github.com/mintutu/restful-golang/model"
)

type PostRepo interface {
	Fetch(ctx context.Context) ([]*model.Post, error)
}

type CommentRepo interface {
	Fetch(ctx context.Context) ([]*model.Comment, error)
}
