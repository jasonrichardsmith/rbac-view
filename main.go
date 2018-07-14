package main

import (
	"log"
	"net/http"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/jasonrichardsmith/rbac-view/client"
	"github.com/jasonrichardsmith/rbac-view/matrix"
	"github.com/jasonrichardsmith/rbac-view/render"
	"github.com/pkg/browser"
)

func main() {

	c, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	m := matrix.New("ClusterRoles")
	log.Print("building permissions matrix")
	err = m.Build(c)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render.Write(m, w)
	})
	browser.OpenURL("http://localhost:8800")
	http.ListenAndServe(":8800", nil)

}
