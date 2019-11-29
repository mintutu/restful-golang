package post

import (
	"context"

	"github.com/mintutu/restful-golang/model"
	"github.com/mintutu/restful-golang/repository"
	"github.com/mintutu/restful-golang/service"
)

const (
	getPostsEndpoint = "https://my-json-server.typicode.com/typicode/demo/posts"
)

type RestPostRepo struct {
	HTTPClient    service.HTTPClient
	RenderHandler service.RenderHandler
}

func NewRestPostRepo() repository.PostRepo {
	return &RestPostRepo{
		HTTPClient:    &service.HTTPClientImpl{},
		RenderHandler: &service.JsonRender{},
	}
}

func (m *RestPostRepo) Fetch(ctx context.Context) ([]*model.Post, error) {
	body, err := m.HTTPClient.GetData(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	posts := make([]*model.Post, 0)
	if err = m.RenderHandler.Render(body, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (m *RestPostRepo) Create(ctx context.Context, p *model.Post) (int64, error) {
	return 0, nil
}
