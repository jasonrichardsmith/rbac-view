package controller

import (
	"errors"

	"github.com/jasonrichardsmith/rbac-view/matrix"
	"github.com/jasonrichardsmith/rbac-view/render"
	"github.com/jasonrichardsmith/rbac-view/render/html"
	"github.com/jasonrichardsmith/rbac-view/render/json"
)

func New(renderer string, matrix matrix.Builder) (render.Renderer, error) {
	switch renderer {
	case "json":
		return json.New(matrix), nil
	case "html":
		return html.New(matrix), nil
	default:
		return nil, errors.New("Inavalid renderer type")
	}
}
