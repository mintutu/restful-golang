package service

import "encoding/json"

type RenderHandler interface {
	Render(data []byte, v interface{}) error
}

type JsonRender struct{}

func (jr *JsonRender) Render(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
