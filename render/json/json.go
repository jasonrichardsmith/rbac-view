package json

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jasonrichardsmith/rbac-view/matrix"
	"github.com/jasonrichardsmith/rbac-view/render"
)

func New(mb matrix.Builder) render.Renderer {
	return &JsonRenderer{
		Writer:  os.Stdout,
		Builder: mb,
	}
}

type JsonRenderer struct {
	Writer  io.Writer
	Builder matrix.Builder
}

func (jr *JsonRenderer) Render() error {
	enc := json.NewEncoder(jr.Writer)
	m, err := jr.Builder.Build()
	if err != nil {
		return err
	}
	return enc.Encode(m)
}
