package view

import (
	"github.com/ericmieth/ub-shift/counter"

	"html/template"
	"net/http"
	"strings"
	"time"

	"database/sql"
)

func BranchAddHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	var e []error
	var err []error

	// POST request
	if r.Method == "POST" {
		inputName := strings.TrimSpace(r.FormValue("name"))
		inputLocation := strings.TrimSpace(r.FormValue("location"))

		e = counter.AddBranch(db, inputName, inputLocation)
		err = append(err, e...)

		if !checkErrors(err) {
			http.Redirect(w, r, "/branch/list/", http.StatusFound)
		}

	}

	// GET request
	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
		Error        []error
	}{
		PageTitle:    "Zweigstelle hinzufügen",
		TemplateName: "/branch/add/",
		StartTime:    startTime,
		Error:        err,
	}

	renderView(w, t, data)
}
