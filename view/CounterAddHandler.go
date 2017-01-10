package view

import (
	"github.com/ericmieth/ub-shift/counter"

	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"database/sql"
)

func CounterAddHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	var err []error

	// POST request
	if r.Method == "POST" {
		inputName := strings.TrimSpace(r.FormValue("name"))
		inputBranch, _ := strconv.Atoi(
			strings.TrimSpace(r.FormValue("branch")),
		)

		e := counter.AddCounter(db, inputName, inputBranch)
		err = append(err, e...)

		if !checkErrors(err) {
			http.Redirect(w, r, "/counter/list/", http.StatusFound)
		}

	}

	// GET request

	branches, e := counter.ListBranches(db)
	err = append(err, e)

	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
		Error        []error
	}{
		PageTitle:    "Theke hinzufügen",
		TemplateName: "/counter/add/",
		ContentData:  branches,
		StartTime:    startTime,
		Error:        err,
	}

	renderView(w, t, data)
}
