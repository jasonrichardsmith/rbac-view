package main

import (
	"log"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/jasonrichardsmith/rbac-view/client"
	"github.com/jasonrichardsmith/rbac-view/matrix"
	"github.com/jasonrichardsmith/rbac-view/render"
	"github.com/jasonrichardsmith/rbac-view/render/controller"
)

func main() {

	c, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	m := matrix.New(c)
	var render render.Renderer
	render, err = controller.New("json", m)
	if err != nil {
		log.Fatal(err)
	}
	err = render.Render()
	if err != nil {
		log.Fatal(err)
	}
}
