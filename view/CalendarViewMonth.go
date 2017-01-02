package view

import (
	"github.com/ericmieth/ub-shift/calendar"

	"html/template"
	"log"
	"net/http"
	"path"
	//"strconv"
	"time"

	"database/sql"
)

func CalendarViewMonthHandler(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	t *template.Template,
) {

	startTime := time.Now()
	base := path.Base(r.URL.Path)

	d, err := time.Parse("2006-01", base)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	year := d.Year()
	month := d.Month()

	days, err := calendar.ReturnMonth(db, year, month)

	// create rows if they dont already exist
	if len(days) == 0 {
		err = calendar.CreateMonth(db, year, month)
		days, err = calendar.ReturnMonth(db, year, month)
		if err != nil {
			log.Print(err)
		}

	}

	content := struct {
		Year  int
		Month time.Month
		Days  []calendar.Day
	}{
		Year:  year,
		Month: month,
		Days:  days,
	}

	data := struct {
		PageTitle    interface{}
		TemplateName string
		ContentData  interface{}
		StartTime    time.Time
	}{
		PageTitle:    "Monatsansicht",
		TemplateName: "/calendar/month/",
		ContentData:  content,
		StartTime:    startTime,
	}

	renderView(w, t, data)
}
