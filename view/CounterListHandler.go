package view

import (
	"github.com/ericmieth/ub-shift/counter"

	"html/template"
	"log"
	"net/http"
	"time"

	"database/sql"
)

func CounterListHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()

	counters, err := counter.ListCounters(db)
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
		PageTitle:    "Theken",
		TemplateName: "/counter/list/",
		ContentData:  counters,
		StartTime:    startTime,
	}

	renderView(w, t, data)
}
