package comment

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mintutu/restful-golang/model"
	"github.com/mintutu/restful-golang/service"
	"github.com/stretchr/testify/assert"
)

func TestCommentFetch(t *testing.T) {
	mockComment := []model.Comment{{ID: 100, Body: "comment test", PostID: 1}}
	CommentBytes, _ := json.Marshal(mockComment)

	tests := map[string]struct {
		expectedComment []*model.Comment
		expectedErr     error
		doHTTPMock      func(httpClientMock *service.MockHTTPClient)
	}{
		"empty Comments": {
			expectedComment: nil,
			expectedErr:     errors.New("unexpected end of JSON input"),
			doHTTPMock: func(httpClientMock *service.MockHTTPClient) {
				httpClientMock.On("GetData", getCommentsEndpoint).Return([]byte{}, nil)
			},
		},
		"parse json properly": {
			expectedComment: []*model.Comment{{ID: 100, Body: "comment test", PostID: 1}},
			expectedErr:     nil,
			doHTTPMock: func(httpClientMock *service.MockHTTPClient) {
				httpClientMock.On("GetData", getCommentsEndpoint).Return(CommentBytes, nil)
			},
		},
	}
	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		t.Run(name, func(t *testing.T) {
			httpMock := &service.MockHTTPClient{}
			renderHandlerMock := &service.JsonRender{}
			test.doHTTPMock(httpMock)

			CommentRepo := RestCommentRepo{
				HTTPClient:    httpMock,
				RenderHandler: renderHandlerMock,
			}
			ctx := context.Background()
			Comments, err := CommentRepo.Fetch(ctx)
			assert.Equal(t, test.expectedComment, Comments)
			if err != nil {
				assert.EqualValues(t, test.expectedErr.Error(), err.Error())
			}
		})
	}
}
