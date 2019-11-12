package comment

import (
	"context"

	"github.com/mintutu/restful-golang/model"
	"github.com/mintutu/restful-golang/repository"
	"github.com/mintutu/restful-golang/service"
)

const getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"

type RestCommentRepo struct {
	HTTPClient    service.HTTPClient
	RenderHandler service.RenderHandler
}

func NewRestCommentRepo() repository.CommentRepo {
	return &RestCommentRepo{
		HTTPClient:    &service.HTTPClientImpl{},
		RenderHandler: &service.JsonRender{},
	}
}

func (m *RestCommentRepo) Fetch(ctx context.Context) ([]*model.Comment, error) {
	body, err := m.HTTPClient.GetData(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	comments := make([]*model.Comment, 0)
	if err = m.RenderHandler.Render(body, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}
