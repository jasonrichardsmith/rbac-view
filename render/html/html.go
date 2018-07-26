package html

import (
	"log"
	"net/http"

	"github.com/jasonrichardsmith/rbac-view/matrix"
	"github.com/jasonrichardsmith/rbac-view/render"
	"github.com/jasonrichardsmith/rbac-view/render/json"
)

func New(mb matrix.Builder) render.Renderer {
	return &HtmlRenderer{
		Builder: mb,
	}
}

type HtmlRenderer struct {
	Builder matrix.Builder
}

func (hr *HtmlRenderer) Render() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index.html", 301)
	})
	http.Handle("/index.html", Handler)
	http.Handle("/dist/", Handler)
	http.Handle("/node_modules/", Handler)
	http.HandleFunc("/allroles.json", func(w http.ResponseWriter, r *http.Request) {
		jr := json.JsonRenderer{
			Builder: hr.Builder,
			Writer:  w,
		}
		err := jr.Render()
		if err != nil {
			log.Println(err)
		}
	})
	return http.ListenAndServe(":8800", nil)

}
