package view

import (
	"github.com/ericmieth/ub-shift/counter"

	"html/template"
	"log"
	"net/http"
	"time"

	"database/sql"
)

func BranchListHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	branches, err := counter.ListBranches(db)
	if err != nil {
		log.Println(err)
	}

	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
		Error        error
	}{
		PageTitle:    "Zweigstellen",
		TemplateName: "/branch/list/",
		ContentData:  branches,
		StartTime:    startTime,
	}

	renderView(w, t, data)
}
