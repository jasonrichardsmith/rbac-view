package render

import (
	"html/template"
	"net/http"

	"github.com/jasonrichardsmith/rbac-view/matrix"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("render/template.html"))
}

func Write(m matrix.Matrix, w http.ResponseWriter) {
	tmpl = template.Must(template.ParseFiles("render/template.html"))
	tmpl.Execute(w, m)
}
