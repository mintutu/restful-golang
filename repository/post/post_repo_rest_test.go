package post

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mintutu/restful-golang/model"
	"github.com/mintutu/restful-golang/service"
	"github.com/stretchr/testify/assert"
)

func TestPostFetch(t *testing.T) {
	mockPost := []model.Post{{ID: 1, Title: "title test"}}
	postBytes, _ := json.Marshal(mockPost)

	tests := map[string]struct {
		expectedPost []*model.Post
		expectedErr  error
		doHTTPMock   func(httpClientMock *service.MockHTTPClient)
	}{
		"empty posts": {
			expectedPost: nil,
			expectedErr:  errors.New("unexpected end of JSON input"),
			doHTTPMock: func(httpClientMock *service.MockHTTPClient) {
				httpClientMock.On("GetData", getPostsEndpoint).Return([]byte{}, nil)
			},
		},
		"parse json properly": {
			expectedPost: []*model.Post{{ID: 1, Title: "title test"}},
			expectedErr:  nil,
			doHTTPMock: func(httpClientMock *service.MockHTTPClient) {
				httpClientMock.On("GetData", getPostsEndpoint).Return(postBytes, nil)
			},
		},
	}
	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		t.Run(name, func(t *testing.T) {
			httpMock := &service.MockHTTPClient{}
			renderHandlerMock := &service.JsonRender{}
			test.doHTTPMock(httpMock)

			postRepo := RestPostRepo{
				HTTPClient:    httpMock,
				RenderHandler: renderHandlerMock,
			}
			ctx := context.Background()
			posts, err := postRepo.Fetch(ctx)
			assert.Equal(t, test.expectedPost, posts)
			if err != nil {
				assert.EqualValues(t, test.expectedErr.Error(), err.Error())
			}
		})
	}
}
