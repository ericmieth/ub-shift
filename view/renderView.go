package view

import (
	"github.com/ericmieth/ub-shift/templates/funcs"

	"bytes"
	"html/template"
	"net/http"
)

func renderView(
	w http.ResponseWriter,
	t *template.Template,
	data interface{},
) {

	funcMap := template.FuncMap{
		"callTemplate": func(
			name string,
			data interface{},
		) (
			template.HTML,
			error,
		) {
			buf := bytes.NewBuffer([]byte{})
			err := t.ExecuteTemplate(buf, name, data)
			if err != nil {
				http.Error(
					w,
					err.Error(),
					http.StatusInternalServerError,
				)
			}
			return template.HTML(buf.String()), err

		},
		"loadTime": funcs.LoadTime,
	}

	err := template.Must(
		template.New("main").
			Funcs(funcMap).
			ParseFiles("templates/main.tmpl")).Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
