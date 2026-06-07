package internal

import (
	"html/template"
	"net/http"
	"path"
)

func Render(w http.ResponseWriter, name string, data any) {
	filePath := path.Join("web", name)

	t := template.Must(template.ParseFiles(
		path.Join("web", "_layout.html"),
		path.Join("web", "_loket_view.html"),
		filePath,
	))
	t.ExecuteTemplate(w, name, data)
}
