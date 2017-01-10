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

func CounterEditHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	var e error
	var err []error

	counterID, e := strconv.Atoi(path.Base(r.URL.Path))
	err = append(err, e)

	// POST request
	if r.Method == "POST" {
		inputName := strings.TrimSpace(r.FormValue("name"))
		inputBranch, _ := strconv.Atoi(
			strings.TrimSpace(r.FormValue("branch")),
		)

		e := counter.EditCounter(db, counterID, inputName, inputBranch)
		err = append(err, e...)

		if !checkErrors(err) {
			http.Redirect(w, r, "/counter/list/", http.StatusFound)
		}

	}

	// GET request

	currentCounter, e := counter.ReturnCounter(db, counterID)
	err = append(err, e)

	branches, e := counter.ListBranches(db)
	err = append(err, e)

	content := struct {
		CounterID   int
		CounterName string
		BranchID    int
		BranchList  []counter.Branch
	}{
		CounterID:   currentCounter.ID,
		CounterName: currentCounter.Name,
		BranchID:    currentCounter.Branch.ID,
		BranchList:  branches,
	}

	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
		Error        []error
	}{
		PageTitle:    "Theke bearbeiten",
		TemplateName: "/counter/edit/",
		ContentData:  content,
		StartTime:    startTime,
		Error:        err,
	}

	renderView(w, t, data)
}
