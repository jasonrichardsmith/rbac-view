package json

import (
	"encoding/json"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
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
	log.Info("Building full matrix for json")
	m, err := jr.Builder.Build()
	log.Info("Matrix for json built")
	if err != nil {
		return err
	}
	return enc.Encode(m)
}
