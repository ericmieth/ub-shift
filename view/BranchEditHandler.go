package view

import (
	"github.com/ericmieth/ub-shift/counter"

	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"database/sql"
)

func BranchEditHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	var e error
	var err []error

	branchID, e := strconv.Atoi(path.Base(r.URL.Path))
	err = append(err, e)

	// POST request
	if r.Method == "POST" {
		inputName := strings.TrimSpace(r.FormValue("name"))
		inputLocation := strings.TrimSpace(r.FormValue("location"))

		err = counter.EditBranch(db, branchID, inputName, inputLocation)
		if !checkErrors(err) {
			http.Redirect(w, r, "/branch/list/", http.StatusFound)
		}

	}

	branch, e := counter.ReturnBranch(db, branchID)
	err = append(err, e)

	// GET request
	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
		Error        []error
	}{
		PageTitle:    "Zweigstelle bearbeiten",
		TemplateName: "/branch/edit/",
		ContentData:  branch,
		StartTime:    startTime,
		Error:        err,
	}

	renderView(w, t, data)
}
