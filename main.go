//go:generate fileb0x fileb0x.yaml
package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/jasonrichardsmith/rbac-view/client"
	"github.com/jasonrichardsmith/rbac-view/matrix"
	"github.com/jasonrichardsmith/rbac-view/render"
	"github.com/jasonrichardsmith/rbac-view/render/controller"
)

var (
	rendertype string
)

func init() {
	flag.StringVar(&rendertype, "render", "html", "render type: json, html")
}

func main() {
	flag.Parse()
	log.Info("Getting K8s client")
	c, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	m := matrix.New(c)
	var render render.Renderer
	render, err = controller.New(rendertype, m)
	if err != nil {
		log.Fatal(err)
	}
	err = render.Render()
	if err != nil {
		log.Fatal(err)
	}
}
