package repository

import (
	"context"

	"github.com/mintutu/restful-golang/model"
)

type PostRepo interface {
	Fetch(ctx context.Context) ([]*model.Post, error)
	Create(ctx context.Context, p *model.Post) (int64, error)
}

type CommentRepo interface {
	Fetch(ctx context.Context) ([]*model.Comment, error)
}
